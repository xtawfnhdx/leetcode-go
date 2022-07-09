package main

import "fmt"

func badCall() {
	panic("bad call")
}

func test13() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("panicing is %s\n", e)
		}
	}()
	badCall()
	fmt.Println("test execute end")
}

func main() {
	fmt.Println("main start")
	test13()
	fmt.Println("main end")

}
