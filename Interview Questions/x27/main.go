package main

import "fmt"

func main() {
	s := make(map[string]int)
	delete(s, "h")
	fmt.Println(s["h"])
}
