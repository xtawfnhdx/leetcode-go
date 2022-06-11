package main

import "fmt"

func main() {
	n := 0
	re := &n
	t1(3, 4, re)
	fmt.Printf("re is %d,n is %d", *re, n)
}
func t1(a, b int, re *int) {
	*re = a * b
}
