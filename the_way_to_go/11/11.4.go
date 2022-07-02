package main

import "fmt"

type Square114 struct {
	test float32
}

func (s *Square114) Area() float32 {
	return s.test
}

type Circle struct {
	test float32
}

func (c Circle) Area() float32 {
	return c.test
}

type IShaper interface {
	Area() float32
}

func main() {
	var ar IShaper
	var sq = Square114{test: 5}
	ar = &sq

	if t, ok := ar.(*Circle); ok {
		fmt.Printf("转换成功 %d\n", t.test)
	} else {
		fmt.Printf("转换失败 %d\n", sq.test)
	}

	if t, ok := ar.(*Square114); ok {
		fmt.Printf("转换成功 %d\n", t.test)
	} else {
		fmt.Printf("转换失败 %d\n", sq.test)
	}
	getType(true)
	getType(Circle{test: 56})
	getType(&Circle{test: 56})

}

func getType(t interface{}) {
	switch t.(type) {
	case bool:
		fmt.Println("this bool")
	case Circle:
		fmt.Println("this Circle")
	case *Circle:
		fmt.Println("this Circle2")
	default:
		fmt.Println("others")
	}
}
