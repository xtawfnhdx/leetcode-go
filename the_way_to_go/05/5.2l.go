package main

import "fmt"

func main() {
	f1()
	f2()
	f3()
}

func f1() {
	a := 3
	switch a {
	case 0:
		fmt.Printf("0")
	case 3:
		fmt.Printf("3")
		fallthrough
	case 5:
		fmt.Printf("5")
	default:
		fmt.Printf("default")
	}
}

func f2() {
	b := 5
	switch {
	case b < 0:
		fmt.Printf("b<0")
	case b > 0:
		fmt.Printf("b>0")
	default:
		fmt.Printf("b=0")
	}
}
func f3() {
	switch c := f3c(); {
	case c < 0:
		fmt.Printf("c<0")
	case c > 0:
		fmt.Printf("c>0")
	default:
		fmt.Printf("c=0")

	}
}
func f3c() int {
	return 5
}
