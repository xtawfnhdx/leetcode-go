package main

import (
	"errors"
	"fmt"
)

type TestError struct{}

func (*TestError) Error() string { return "" }

func main() {

	// 错误链。
	a := errors.New("a")
	b := fmt.Errorf("b, %w", a)
	c := fmt.Errorf("c, %w", b)

	fmt.Println(c)                     // c, b, a
	fmt.Println(errors.Unwrap(c) == b) // true

	// 递归检查。
	fmt.Println(errors.Is(c, a)) // true

	// -------------------------------

	x := &TestError{}
	fmt.Println(&x)
	y := fmt.Errorf("y, %w", x)
	z := fmt.Errorf("z, %w", y)

	// 二级指针，写入错误对象地址。
	var x2 *TestError
	if errors.As(z, &x2) {
		fmt.Println(x == x2) // true
	}
}
