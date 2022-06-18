package main

import "fmt"

func main() {
	var arr1 [5]int
	for i := 0; i < 5; i++ {
		arr1[i] = i * i
	}
	for i := 0; i < 5; i++ {
		fmt.Printf("arr%d res is %d\n", i, arr1[i])
	}
	for i, res := range arr1 {
		fmt.Printf("arr%d res is %d\n", i, res)

	}
}
