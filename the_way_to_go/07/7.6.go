package main

import "fmt"

func Sum(a *[10]int) int {
	res := 0
	for _, x := range a {
		res += x
	}
	return res
}

func main() {
	arr1 := [10]int{3, 4, 5, 6, 7, 8, 9, 12}
	fmt.Println(arr1)
	fmt.Printf("res is %d\n", Sum(&arr1))
}
