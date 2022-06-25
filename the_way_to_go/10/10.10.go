package main

import "fmt"

type TwoInts struct {
	a, b int
}

func main() {
	two := TwoInts{3, 4}
	fmt.Println(two.Add())
	fmt.Println(two.AddParam(5))
}
func (tn *TwoInts) Add() int {
	return tn.a + tn.b
}
func (tn *TwoInts) AddParam(param int) int {
	return tn.a + tn.b + param
}
