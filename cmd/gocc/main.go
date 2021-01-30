package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/teamlint/opencc"
)

var (
	input  = flag.String("input", "", "file of original text to read")
	output = flag.String("output", "", "file of converted text to write")
	config = flag.String("config", "", "convert config, s2t, t2s, etc")
)

func main() {
	flag.Parse()
	var err error
	var in, out *os.File //io.Reader
	if *input == "" {
		in = os.Stdin
	} else {
		in, err = os.Open(*input)
		if err != nil {
			log.Fatal(err)
		}
		defer in.Close()
	}
	br := bufio.NewReader(in)

	if *output == "" {
		out = os.Stdout
	} else {
		out, err = os.OpenFile(*output, os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
	}

	if *config == "" {
		*config = "s2t"
	}

	conv, err := opencc.New(*config)
	if err != nil {
		log.Fatal(err)
	}

	err = forEachLine(br, func(line string) error {
		str, e := conv.Convert(line)
		if e != nil {
			return e
		}
		fmt.Fprint(out, str+"\n")
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}

func forEachLine(br *bufio.Reader, callback func(string) error) error {
	stop := false
	for {
		if stop {
			break
		}
		line, err := br.ReadString('\n')
		if err == io.EOF {
			stop = true
		} else if err != nil {
			return err
		}
		line = strings.TrimSuffix(line, "\n")
		if line == "" {
			if !stop {
				if err = callback(line); err != nil {
					return err
				}
			}
			continue
		}
		if err = callback(line); err != nil {
			return err
		}
	}
	return nil
}
