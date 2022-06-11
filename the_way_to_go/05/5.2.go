package main

import "fmt"

func main() {
	first := 10
	var cond int
	if first <= 0 {
		fmt.Printf("first<=0")
	} else if first > 0 && first < 5 {
		fmt.Printf("<0first<5")
	} else {
		fmt.Printf("first>=5")
	}
	if cond = 10; cond > 10 {
		fmt.Printf("cond>10")
	} else {
		fmt.Printf("fmt<=10")
	}
}
