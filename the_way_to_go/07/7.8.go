package main

import "fmt"

func main() {
	slice1 := make([]int, 10)
	for i := 0; i < len(slice1); i++ {
		slice1[i] = 5 * i
	}
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("slice len is %d cap is %d", len(slice1), cap(slice1))
}
