package main

import (
	"fmt"
)

func main(){
  s := make([]int, 10)

  s = append(s, 1, 2, 3)

  fmt.Println(s)
}