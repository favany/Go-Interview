package main

import (
	"fmt"
	"strings"
)

func main() {
	re1 := reverseString("abcde")
	fmt.Println(re1)
	re2 := reverseString("123")
	fmt.Println(re2)
}

func reverseString(str string) string {
	if len(str) > 5000 {
		return "length over 5000, error."
	}
	
	s := strings.Split(str, "")
	for i, j := 0, len(str) - 1; i < j; i, j = i + 1, j - 1 {
		s[i], s[j] = s[j], s[i]
	}
	return strings.Join(s, "")
}