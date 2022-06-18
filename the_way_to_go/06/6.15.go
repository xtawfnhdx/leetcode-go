package main

import "fmt"

func main() {
	f()
	fmt.Println(f2())
}
func f() {
	for i := 0; i < 5; i++ {
		g := func(i int) { fmt.Printf("i=%d\n", i) }
		g(i)
		fmt.Printf("g is type %T and vaule is %v\n", g, g)
	}
}
func f2() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}
