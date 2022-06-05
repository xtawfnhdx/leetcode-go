package main

import "fmt"

func main() {
	capitals := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}
	for key := range capitals {
		fmt.Println(key)
	}
}
func maintest2() {
	slt1 := []int{1, 2, 3}
	slt2 := append(slt1, 4, 5, 6)
	fmt.Println("slt2=", slt2)
	slt3 := append(slt1, slt2...)
	fmt.Println("slt3=", slt3)
}

func maintest() {
	slFrom := []int{1, 2, 3}
	slTo := make([]int, 10)
	fmt.Println("from=", slFrom, "to=", slTo)
	//from= [1 2 3] to= [0 0 0 0 0 0 0 0 0 0]
	copy(slTo, slFrom)
	fmt.Println("from=", slFrom, "to=", slTo)
	//from= [1 2 3] to= [1 2 3 0 0 0 0 0 0 0]
	slFrom[1] = 999
	fmt.Println("from=", slFrom, "to=", slTo)
	//from= [1 999 3] to= [1 2 3 0 0 0 0 0 0 0]
	slFrom2 := []int{1, 2, 3, 4, 5}
	slTo2 := make([]int, 3)
	cou := copy(slTo2, slFrom2)
	fmt.Println("from=", slFrom2, "to=", slTo2, "cou=", cou)
	//from= [1 2 3 4 5] to= [1 2 3] cou= 3
}
