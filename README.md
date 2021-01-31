# Golang OpenCC
Golang 简体繁体中文互转

[![Go Report Card](https://goreportcard.com/badge/github.com/teamlint/opencc)](https://goreportcard.com/report/github.com/teamlint/opencc) [![GoDoc](https://godoc.org/github.com/teamlint/opencc?status.svg)](https://godoc.org/github.com/teamlint/opencc) [![GitHub release](https://img.shields.io/github/release/teamlint/opencc.svg)](https://github.com/teamlint/opencc/releases/latest)

## Introduction 介紹
gocc is a golang port of OpenCC([Open Chinese Convert 開放中文轉換](https://github.com/BYVoid/OpenCC/)) which is a project for conversion between Traditional and Simplified Chinese developed by [BYVoid](https://www.byvoid.com/).

gocc stands for "**Go**lang version Open**CC**", it is a total rewrite version of OpenCC in Go. It just borrows the dict files and config files of OpenCC, so it may not produce the same output with the original OpenCC.

 参考以下两个仓库源进行优化,并使用Go Module进行管理, 方便新项目引用, 同时进行完善更新
> [gocc](https://github.com/liuzl/gocc)
> [OpenCC-go](https://github.com/ApesPlan/OpenCC-go)



## Installation 安裝

### 1. Golang Package
```sh
go get github.com/teamlint/opencc
```

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/teamlint/opencc"
)

func main() {
    // 简体转繁体
    s2t, err := openc.New("s2t")
    if err != nil {
        log.Fatal(err)
    }
    in := `自然语言处理是人工智能领域中的一个重要方向。`
    out, err := s2t.Convert(in)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s\n%s\n", in, out)
    //自然语言处理是人工智能领域中的一个重要方向。
    //自然語言處理是人工智能領域中的一個重要方向。

    // 繁体转简体
    t2s, err := ccgo.New("t2s")
    if err != nil {
        log.Fatal(err)
    }
    in := "閱坊-閱讀的樂趣"
    out, err := t2s.Convert(str)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s\n%s\n", in, out)
    //閱坊-閱讀的樂趣
    //阅坊-阅读的乐趣
}
```
### 2. Command Line
```sh
git clone https://github.com/teamlint/opencc
cd opencc/cmd/gocc
make install
gocc --help
echo "阅坊-阅读的乐趣" | gocc
#閱坊-閱讀的樂趣
```



## Conversions 语言转换

**目前支持14种**

s2t, t2s, s2tw, tw2s, s2hk, hk2s, s2twp, tw2sp, t2tw, hk2t, t2hk, t2jp, jp2t, tw2t

1. s2t ==> Simplified Chinese to Traditional Chinese 簡體到繁體
2. t2s ==> Traditional Chinese to Simplified Chinese 繁體到簡體
3. s2tw ==> Simplified Chinese to Traditional Chinese (Taiwan Standard) 簡體到臺灣正體
4. tw2s ==> Traditional Chinese (Taiwan Standard) to Simplified Chinese 臺灣正體到簡體
5. s2hk ==> Simplified Chinese to Traditional Chinese (Hong Kong variant) 簡體到香港繁體
6. hk2s ==> Traditional Chinese (Hong Kong variant) to Simplified Chinese 香港繁體到簡體
7. s2twp ==> Simplified Chinese to Traditional Chinese (Taiwan Standard) with Taiwanese idiom 簡體到繁體（臺灣正體標準）並轉換爲臺灣常用詞彙
8. tw2sp ==> Traditional Chinese (Taiwan Standard) to Simplified Chinese with Mainland Chinese idiom 繁體（臺灣正體標準）到簡體並轉換爲中國大陸常用詞彙
9. t2tw ==> Traditional Chinese (OpenCC Standard) to Taiwan Standard 繁體（OpenCC 標準）到臺灣正體
10. hk2t ==> Traditional Chinese (Hong Kong variant) to Traditional Chinese 香港繁體到繁體（OpenCC 標準）
11. t2hk ==> Traditional Chinese (OpenCC Standard) to Hong Kong variant 繁體（OpenCC 標準）到香港繁體
12. t2jp ==> Traditional Chinese Characters (Kyūjitai) to New Japanese Kanji (Shinjitai) 繁體（OpenCC 標準，舊字體）到日文新字體
13. jp2t ==> New Japanese Kanji (Shinjitai) to Traditional Chinese Characters (Kyūjitai) 日文新字體到繁體（OpenCC 標準，舊字體）
14. tw2t ==> Traditional Chinese (Taiwan standard) to Traditional Chinese 臺灣正體到繁體（OpenCC 標準）



## Updates 使用最新词典

### **更新词典文件**

获取[最新OpenCC代码](https://github.com/BYVoid/OpenCC)
使用 `OpenCC/data/config/*.json` 和 `OpenCC/data/dictionary/*.txt` 
替换本包的 `config/*.json` 和 `dictionary/*.txt` 相关文件

`OpenCC/data/config/*.json`文件中 默认匹配的是`.ocd2`文件（`"type": "ocd2", "file": "TSPhrases.ocd2"）`，全部替换为`txt`即可

### 生成最新词典文件

以下文件由部分词典文件进一步操作产生, 需要手动处理或使用 [OpenCC 脚本处理](https://github.com/BYVoid/OpenCC/tree/master/data/scripts)

- `HKVariantsRev.txt` 由 `HKVariants.txt` 反转列产生
- `JPVariantsRev.txt` 由 `JPVariants.txt` 反转列产生
- `TWPhrases.txt` 由 `TWPhrasesIT.txt` `TWPhrasesName.txt` `TWPhrasesOther.txt` 合并产生
- `TWPhrasesRev.txt` 由 `TWPhrases.txt` 反转列产生
- `TWVariantsRev.txt` 由 `TWVariants.txt` 反转列产生

### 添加新的语言包

如果有新添加的语言, 修改 `opencc.go` 文件中 `supportedConversions` `conversions` 值, 同时增加相关词典文件即可：

```go
	supportedConversions = "s2t, t2s, s2tw, tw2s, s2hk, hk2s, s2twp, tw2sp, t2tw, hk2t, t2hk, t2jp, jp2t, tw2t"
    ......
	conversions          = map[string]struct{}{
		"s2t":   {},
		"t2s":   {},
		"s2tw":  {},
		"tw2s":  {},
		"s2hk":  {},
		"hk2s":  {},
		"s2twp": {},
		"tw2sp": {},
		"t2tw":  {},
		"hk2t":  {},
		"t2hk":  {},
		"t2jp":  {},
		"jp2t":  {},
		"tw2t":  {},
	}
```

