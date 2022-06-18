package main

import "fmt"

func main() {
	r := f17()
	fmt.Printf("r(1) is %d\n", r(1))
	fmt.Printf("r(1) is %d\n", r(10))
	fmt.Printf("r(1) is %d\n", r(100))

}
func f17() func(i int) int {
	var x int
	return func(i int) int {
		x += i
		return x
	}
}
