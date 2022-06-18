package main

import "fmt"

func main() {
	p := Add2()
	fmt.Printf("add2 res is %d\n", p(3))
	w := Add1(3)
	fmt.Printf("add1 res is %d\n", w(10))

}
func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Add1(i int) func(b int) int {
	return func(b int) int {
		return i + b
	}
}
