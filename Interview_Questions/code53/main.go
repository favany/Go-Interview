package main

import (
	"fmt"
)

type S struct {
    m string
}

func f() *S {
    return &S{}  //A
}

func main() {
  p := S{m:"foo"}   //B
  fmt.Println(p.m) //print "foo"
}