package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	re := checkConformityAfterSort("aabcd", "dbace")
	fmt.Println(re)
}

func checkConformityAfterSort(str1, str2 string) bool {
	return sortStr(str1) == sortStr(str2)
}

func sortStr(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}