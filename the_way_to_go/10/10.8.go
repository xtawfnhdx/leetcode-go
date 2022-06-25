package main

import "fmt"

type innerS struct {
	in1 int
	in2 int
}
type outreS struct {
	b int
	int
	innerS
}

func main() {
	out := new(outreS)
	out.b = 3
	out.int = 5
	out.in1 = 4
	out.in2 = 2
	fmt.Println(out)

	out2 := outreS{2, 3, innerS{5, 6}}
	fmt.Println(out2)
}
