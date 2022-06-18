package main

import "fmt"

func main() {
	fmt.Printf("10 的斐波那契数列为：%d\n", fibnac(10))
}

func fibnac(n int) int {
	if n <= 1 {
		return 1
	} else {
		return fibnac(n-1) + fibnac(n-2)
	}
}
