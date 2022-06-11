package main

import "fmt"

func main() {
	re1, re2 := fts1(2, 3, 4)
	fmt.Printf("re1 is %d re2 is %d \n", re1, re2)
	rs1, rs2 := fts2(2, 3, 4)
	fmt.Printf("rs1 is %d rs2 is %d \n", rs1, rs2)
}
func fts1(a, b, c int) (int, int) {
	r1 := a + b + c
	r2 := a * b * c
	return r1, r2
}

func fts2(a, b, c int) (x1, x2 int) {
	x1 = a + b + c
	x2 = a * b * c
	return

}
