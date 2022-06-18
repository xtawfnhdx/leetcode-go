package main

import "fmt"

func main() {
	a()
}
func funcin(funcname string) {
	fmt.Printf("in func %s\n", funcname)
}
func funcout(funcname string) {
	fmt.Printf("out func %s\n", funcname)
}

func a() {
	funcin("a")
	defer funcout("a")
	fmt.Println("this func a action")
	b()
}
func b() {
	funcin("b")
	defer funcout("b")
	fmt.Println("this func b action")
}
