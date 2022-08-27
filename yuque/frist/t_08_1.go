package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main2() {
	go func() {
		defer fmt.Println("go done")
		time.Sleep(time.Second)
	}()
	defer fmt.Println("main done")
}
func main3() {
	q := make(chan struct{})
	go func() {
		defer close(q)
		fmt.Println("go done")
		time.Sleep(time.Second)
	}()
	<-q
	defer fmt.Println("main done")
}
func main4() {
	var ws sync.WaitGroup
	ws.Add(10)
	for i := 0; i < 10; i++ {
		go func(x int) {
			defer ws.Done()
			fmt.Println(x, " done")
		}(i)
	}
	ws.Wait()
	fmt.Println("main done")
}

func main5() {
	q := make(chan struct{})
	go func() {
		defer close(q)
		defer fmt.Println("go done")
		a()
		b()
		c()
	}()
	<-q
	fmt.Println("main done")
}
func a() { fmt.Println("a") }
func b() { fmt.Println("b"); runtime.Goexit() }
func c() { fmt.Println("c") }

func main() {
	quit := make(chan struct{})
	data := make(chan int)
	go func() {
		data <- 11
	}()
	go func() {
		defer fmt.Println("go done 2")
		//defer close(quit)
		defer fmt.Println("go done")
		fmt.Println(<-data)
		fmt.Println(<-data)
	}()
	data <- 12
	fmt.Println("等待")
	<-quit
	fmt.Println("结束")
}
