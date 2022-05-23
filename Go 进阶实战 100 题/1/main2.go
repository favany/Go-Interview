package main

import (
	"fmt"
	"sync"
	"unicode/utf8"
)

func main() {
        number, letter := make(chan bool), make(chan bool)
        wait := sync.WaitGroup{}
        go func() {
					i := 1
					for {
						<- number
						fmt.Printf("%d%d", i, i + 1)
						i += 2
						letter <- true
					}
        }()
        wait.Add(1)
        go func(wait *sync.WaitGroup) {
					str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
					i := 0
					for {
						<- letter
						if i >= utf8.RuneCountInString(str) {
										wait.Done()
										return
						}
						fmt.Print(str[i : i+2])
						i += 2
						number <- true
					}
        }(&wait)
        number <- true
        wait.Wait()
}

