package main

import (
	"fmt"
	"time"
)

var resume chan int

func integers() chan int {
	yield := make(chan int)
	count := 0
	go func() {
		for {
			fmt.Println("写入管道")
			time.Sleep(1e9)
			yield <- count
			fmt.Println("写入管道完成")
			count++
		}
	}()
	return yield
}

func geterateInterage() int {
	fmt.Println("准备获取数据")
	time.Sleep(1e9)
	x := <-resume
	time.Sleep(1e9)
	fmt.Println("已经获取了数据")
	return x
}
func main() {
	resume = integers()
	fmt.Println(geterateInterage())
	//fmt.Println(geterateInterage())
	//fmt.Println(geterateInterage())
}
