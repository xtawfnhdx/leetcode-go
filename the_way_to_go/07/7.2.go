package main

import "fmt"

func f(a [3]int) {
	fmt.Println(a)
}
func fp(a *[3]int) {
	fmt.Println(a)
}

func main() {
	var ar [3]int
	f(ar)
	fp(&ar)

	var arr1 = new([5]int)
	var arr2 [5]int
	fmt.Printf("arr1 type is %T\n", arr1)
	fmt.Printf("arr2 type is %T\n", arr2)
	arr3 := arr1
	arr4 := *arr1
	fmt.Println(arr1)
	fmt.Println(arr3)
	fmt.Println(arr4)
	arr3[2] = 10
	arr4[3] = 20
	fmt.Println(arr1)
	fmt.Println(arr3)
	fmt.Println(arr4)
}
