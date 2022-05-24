package main

import (
	"fmt"
	"strings"
	"unicode"
)

func replaceSpace(str string) (string, bool) {
	s := strings.Split(str, "")
	if len(s) > 1000 {
		return str, false
	}
	// unicode.IsLetter 判断字符是否是字⺟
	for _, v := range str {
		if unicode.IsLetter(v) == false && string(v) != " " {
			return str, false
		}
	}
	return strings.Replace(str, " ", "%20", -1), true
}

func main() {
	re, ok := replaceSpace("aaccc xx")
	if ok {
		fmt.Println(re)
	} else {
		fmt.Println("不符合条件")
	}

}