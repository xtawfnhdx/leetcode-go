package main

import "fmt"

type A struct {
	ax, ay int
}
type B struct {
	A
	bx, by int
}

func main() {
	b := B{A{1, 2}, 3, 4}
	fmt.Println(b)
	fmt.Println(b.A)
	fmt.Println(b.ax, b.ay)
}
