package main

import (
	"fmt"
)

type S struct {
    m string
}

func f() *S {
    return "foo"  //A
}

func main() {
  p := S{m:f}   //B
    fmt.Println(p.m) //print "foo"
}