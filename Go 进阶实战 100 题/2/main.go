package main

import "fmt"

func main() {
	re1 := checkString("aaabb")
	fmt.Println(re1)
	re2 := checkString("abcde")
	fmt.Println(re2)
}

func checkString(str string) bool {
	for i1, v1 := range str {
		for i2, v2 := range str {
			if i1 != i2 && v1 == v2 {
				return false
			}
		}
	}
	return true
}