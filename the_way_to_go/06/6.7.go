package main

import "fmt"

func main() {
	x := []int{4, 6, 3}
	r1 := f1(x...)
	fmt.Printf("min is %d\n", r1)

	r2 := f1(3, 5, 7, 2, 5)
	fmt.Printf("r2 is %d \n", r2)

	ft1(3, 4, 5)
	ft4(3, 4, Options{pam1: 6, pam3: false, pam2: "abc"})
	ft5(3, 4, 34, true, "aaa", 3.45)
}
func f1(par ...int) int {
	if len(par) == 0 {
		return 0
	}
	re := par[0]
	for _, r := range par {
		if r <= re {
			re = r
		}
	}
	return re
}

func ft1(pam ...int) {
	ft2(pam...)
	ft3(pam)
}
func ft2(pam ...int) {
	for _, r := range pam {
		fmt.Printf("ft2:%d\n", r)
	}
}
func ft3(pam []int) {
	for _, r := range pam {
		fmt.Printf("ft3:%d\n", r)
	}
}

func ft4(a int, b int, options Options) {
	fmt.Printf("a is %d\n", a)
	fmt.Printf("b is %d\n", b)
	fmt.Printf("op.p1 is %d\n", options.pam1)
	fmt.Printf("op.p2 is %s\n", options.pam2)
	fmt.Printf("op.p3 is %t\n", options.pam3)

}

type Options struct {
	pam1 int
	pam2 string
	pam3 bool
}

func ft5(a int, b int, values ...interface{}) {
	fmt.Printf("a is %d\n", a)
	fmt.Printf("b is %d\n", b)
	for _, x := range values {
		switch v := x.(type) {
		case int:
			fmt.Printf("v is %d\n", v)
		case string:
			fmt.Printf("v is %s\n", v)
		case bool:
			fmt.Printf("v is %t\n", v)
		default:
			fmt.Print(v)

		}
	}
}
