package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (s Square) Area() float32 {
	return s.side * s.side
}
func main() {
	sq1 := new(Square)
	sq1.side = 5
	var ar Shaper
	ar = sq1
	fmt.Printf("res is %f", ar.Area())
}
