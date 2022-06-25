package main

import "fmt"

type ListTest []int

func (ll ListTest) Sum() (re int) {
	for _, v := range ll {
		re += v
	}
	return
}

func main() {
	fmt.Println(ListTest{1, 2, 3, 4}.Sum())
}
