package opencc

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-ego/cedar"
)

// Dict contains the Trie and dict values
type Dict struct {
	Trie   *cedar.Cedar
	Values [][]string
}

// BuildFromFile builds the da dict from fileName
// BuildFromFile从fileName生成数据
func BuildFromFile(fileName string) (*Dict, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return Build(file)
}

// Build da dict from io.Reader
// 从io.Reader构建数据
func Build(in io.Reader) (*Dict, error) {
	trie := cedar.New()
	values := [][]string{}
	br := bufio.NewReader(in)
	for {
		line, err := br.ReadString('\n')
		if err == io.EOF {
			break
		}
		items := strings.Split(strings.TrimSpace(line), "\t")
		if len(items) < 2 {
			continue
		}
		err = trie.Insert([]byte(items[0]), len(values))
		if err != nil {
			return nil, err
		}

		if len(items) > 2 {
			values = append(values, items[1:])
		} else {
			values = append(values, strings.Fields(items[1]))
		}
	}
	return &Dict{Trie: trie, Values: values}, nil
}

// Load gob serialized dict from dir
// 从目录加载gob序列化的dict
func Load(dir string) (*Dict, error) {
	trieFile := filepath.Join(dir, "trie")
	valueFile := filepath.Join(dir, "values")
	trie := cedar.New()
	if err := trie.LoadFromFile(trieFile, "gob"); err != nil {
		return nil, err
	}

	file, err := os.OpenFile(valueFile, os.O_RDONLY, 0600)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	in := bufio.NewReader(file)
	dataDecoder := gob.NewDecoder(in)
	var values [][]string
	if err = dataDecoder.Decode(&values); err != nil {
		return nil, err
	}
	return &Dict{Trie: trie, Values: values}, nil
}

// PrefixMatch str by Dict, returns the matched string and its according values
// Dict的PrefixMatch str，返回匹配的字符串及其相应的值
func (d *Dict) PrefixMatch(str string) (map[string][]string, error) {
	if d.Trie == nil {
		return nil, fmt.Errorf("Trie is nil")
	}
	ret := make(map[string][]string)
	for _, id := range d.Trie.PrefixMatch([]byte(str), 0) {
		key, err := d.Trie.Key(id)
		if err != nil {
			return nil, err
		}
		value, err := d.Trie.Value(id)
		if err != nil {
			return nil, err
		}
		ret[string(key)] = d.Values[value]
	}
	return ret, nil
}

// Get the values of str, like map
// 获取str的值，例如map
func (d *Dict) Get(str string) ([]string, error) {
	if d.Trie == nil {
		return nil, fmt.Errorf("trie is nil")
	}
	id, err := d.Trie.Get([]byte(str))
	if err != nil {
		return nil, err
	}
	value, err := d.Trie.Value(id)
	if err != nil {
		return nil, err
	}
	return d.Values[value], nil
}

// Save gob serialized dict to dir
// 保存Gob序列化字典到目录
func (d *Dict) Save(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0755)
	}
	trieFile := filepath.Join(dir, "trie")
	valueFile := filepath.Join(dir, "values")
	if err := d.Trie.SaveToFile(trieFile, "gob"); err != nil {
		return err
	}
	file, err := os.OpenFile(valueFile, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	out := bufio.NewWriter(file)
	defer out.Flush()
	dataEncoder := gob.NewEncoder(out)
	return dataEncoder.Encode(d.Values)
}
