package main

import "fmt"

func parseString(str string) string {
	len := len(str)
	s := make([]string, len)
	recordSlice := make([]string, len)
	sign := false
	for i := len - 1; i >= 0; i++ {
		if string(str[i]) == ")" {
			sign = true
		} else if string(str[i]) == "(" {
			sign = false
		}
		if sign == true {
			recordSlice = append(recordSlice, string(str[i]))
		} else if sign == false {
			s = append(s, recordSlice...)
			recordSlice = make([]string, len)
			
		}

	}
}

func getPosition(str string) (int, int) {
	var x int
	var y int
	for _, v := range str {
		switch string(v) {
		case "R":
			x ++
		case "L":
			x --
		case "F":
			y ++
		case "B":
			x --
		}
	}
	return x, y
}

func main() {
	// fmt.Println(getPosition("RLFLF"))
	fmt.Println(getPosition("12(LRR)"))
}