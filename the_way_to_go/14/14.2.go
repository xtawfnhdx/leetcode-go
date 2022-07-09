package main

import (
	"fmt"
	"time"
)

func main() {
	//var ch chan string
	//ch = make(chan string)
	ch := make(chan string)
	go send(ch)
	go getstr(ch)
	time.Sleep(10 * 1e9)
}

func send(ch chan string) {
	ch <- "first"
	ch <- "second"
	ch <- "third"
	ch <- "fore"
}
func getstr(ch chan string) {
	time.Sleep(2 * 1e9)

	var input string
	for {
		input = <-ch
		fmt.Printf("%s ", input)
	}
}
