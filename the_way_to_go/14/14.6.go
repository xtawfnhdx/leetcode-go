package main

import (
	"fmt"
	"time"
)

func main() {
	sump2(crch())
	time.Sleep(1e8)
}
func crch() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func sump2(ch chan int) {
	go func() {
		for x := range ch {
			fmt.Println(x)
		}
	}()
}
