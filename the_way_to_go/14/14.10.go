package main

import (
	"fmt"
	"time"
)

func pump141(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}
func pump142(ch chan int) {
	for i := 0; ; i++ {
		ch <- i + 5
	}
}
func suck14(ch1 chan int, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Println("ch1 ", v)
		case v := <-ch2:
			fmt.Println("ch2 ", v)
		}
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go pump141(ch1)
	go pump141(ch2)
	go suck14(ch1, ch2)
	time.Sleep(1e9)
}
