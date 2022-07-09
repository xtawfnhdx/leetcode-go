package main

import (
	"fmt"
	"time"
)

func main() {
	ch := sump()
	go unluck(ch) //不加go  就变成死循环咯
	time.Sleep(1e8)
}
func sump() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func unluck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
