package main

import "fmt"

type te1 struct {
	name string
	age  int
}

func main() {
	var t1 te1
	t1.name = "zhangsan"
	t1.age = 12

	var t2 *te1 = new(te1)
	t2.name = "lisi"
	t2.age = 13

	t3 := new(te1)
	t3.name = "wangwu"
	t3.age = 14

	t4 := &te1{name: "zhaoliu", age: 16}
	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(t3)
	fmt.Println(t4)

}
