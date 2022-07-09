package main

import (
	"fmt"
	"time"
)

func f1(in chan int) {
	fmt.Println(<-in)
}
func main() {
	//无缓冲通道，会导致死锁
	//in := make(chan int)
	//有缓冲通道，可以执行成功
	in := make(chan int, 10)
	in <- 2
	go f1(in)
	time.Sleep(1e9)
}
