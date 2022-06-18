package main

import "fmt"

func main() {
	a3()
}
func funcin(funcname string) {
	fmt.Printf("in func %s\n", funcname)
}
func funcout(funcname string) {
	fmt.Printf("out func %s\n", funcname)
}

func a3() {
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
