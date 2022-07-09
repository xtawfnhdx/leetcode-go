package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main start")
	go longWait()
	go shortWait()
	fmt.Println("main end")
	time.Sleep(10 * 1e9)
	fmt.Println("main sleep end")
}
func longWait() {
	fmt.Println("begin logwait")
	time.Sleep(5 * 1e9)
	fmt.Println("end logwait")
}

func shortWait() {
	fmt.Println("begin short")
	time.Sleep(2 * 1e9)
	fmt.Println("end short")
}
