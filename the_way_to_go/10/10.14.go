package main

import "fmt"

type List2 []int

func (l List2) Len() int {
	return len(l)
}

func (l *List2) Append(val int) {
	*l = append(*l, val)
}

func main() {
	var ls1 List2
	ls1.Append(12)
	fmt.Println(ls1, ls1.Len())

	ls2 := new(List2)
	ls2.Append(23)
	fmt.Println(ls2, ls2.Len())
}