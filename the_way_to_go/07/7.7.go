package main

import "fmt"

func main() {
	var arr1 [6]int
	var slice1 []int = arr1[2:5]
	for i := 0; i < 5; i++ {
		arr1[i] = i
	}
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("slice %d res is %d\n", i, slice1[i])
	}

	fmt.Printf("arr1 len is %d\n", len(arr1))
	fmt.Printf("slice1 len is %d\n", len(slice1))
	fmt.Printf("slice1 cap is %d\n", cap(slice1))

	slice1 = slice1[0:4]

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("slice %d res is %d\n", i, slice1[i])
	}
	fmt.Printf("slice1 len is %d\n", len(slice1))
	fmt.Printf("slice1 cap is %d\n", cap(slice1))
}
