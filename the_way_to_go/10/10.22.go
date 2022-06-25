package main

import (
	"fmt"
	"strconv"
)

type TwoInts2 struct {
	a, b int
}

func (t *TwoInts2) String() string {
	return "(" + strconv.Itoa(t.a) + "/" + strconv.Itoa(t.b) + ")"
}
func main() {
	tre := new(TwoInts2)
	tre.a = 3
	tre.b = 4
	fmt.Printf("re is %v \n", tre)
	fmt.Println(tre)
	fmt.Printf("re is %T \n", tre)
	fmt.Printf("re is %#v \n", tre)
}
