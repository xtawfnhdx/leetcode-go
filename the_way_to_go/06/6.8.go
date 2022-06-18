package main

import "fmt"

func main() {
	function1()
	a()
}

func function1() {
	fmt.Printf("In function1 at the top\n")
	defer function2()
	fmt.Printf("In function1 at the bottom \n")
}

func function2() {
	fmt.Printf("this function action res\n")
}

func a() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("i=%d \n", i)
	}
}
