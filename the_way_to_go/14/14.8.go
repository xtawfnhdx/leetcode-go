package main

import "fmt"

func generate() chan int {
	in := make(chan int)
	go func() {
		for i := 2; i < 100; i++ {
			in <- i
		}
	}()
	return in
}

func filter(in chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func sieve() chan int {
	out := make(chan int)
	go func() {
		ch := generate()
		for {
			prime := <-ch
			ch = filter(ch, prime)
			out <- prime
		}
	}()
	return out
}
func main() {
	pr := sieve()
	for {
		fmt.Println(<-pr)
	}
}
