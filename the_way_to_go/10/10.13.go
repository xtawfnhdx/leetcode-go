package main

import "fmt"

type B102 struct {
	thing int
}

func (b *B102) change() {
	b.thing = 1
}
func (b B102) write() string {
	return fmt.Sprint(b)
}

func main() {
	//值类型
	var b B102
	b.change()
	fmt.Println(b.write())

	//引用类型
	b2 := B102{}
	b2.change()
	fmt.Println(b2.write())
}
