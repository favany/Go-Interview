package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	numCh := make(chan int)
	runeCh := make(chan rune)
	wg.Add(1)

	go func () {
		var num int
		for i := 1; i <= 14; {
			num = <- numCh
			fmt.Print(num + 1)
			fmt.Print(num + 2)
			i++
			if i <= 14 {
				runeCh <- rune(num)
			}
		}
		wg.Done()
	}()

	go func () {
		var run rune
		for i := 1; i <= 13; {
			run = <- runeCh + 65
			fmt.Print(string(run))
			fmt.Print(string(run+1))
			i++
			numCh <- int(run) + 2 - 65
		}
	}()

	numCh <- 0
	wg.Wait()
}