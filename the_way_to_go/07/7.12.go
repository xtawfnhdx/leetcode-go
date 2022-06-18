package main

import "fmt"

func main() {
	slForm := []int{1, 2, 3}
	slTo := make([]int, 10)
	n := copy(slTo, slForm)
	fmt.Println(n)
	fmt.Println(slTo)
	sl3 := []int{1, 2, 4}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)

	sl4 := []int{1, 2, 3}
	sl5 := []int{4, 5, 6}
	sl4 = append(sl4, sl5...)
	fmt.Println(sl4)

}
