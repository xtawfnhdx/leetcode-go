package main

import "fmt"

const c = "C"

var v int = 5

type T struct {
}

//init 该函数不需要手动调用，会自动调用
func init() {
	fmt.Print("init action")
}

func main() {
	var a int = 6
	Func1()
	fmt.Print(a, c, v, new(T))
}
func Func1() {

}
