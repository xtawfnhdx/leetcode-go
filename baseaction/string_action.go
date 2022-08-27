package baseaction

import (
	"fmt"
)

const (
	TestName  = "zhangsan"
	TestName2 = "李四"
)

func ShowConstString() {
	fmt.Println("TestName:", TestName)
	fmt.Println("TestName2:", TestName2)
}

func ShowForEach() {
	t1 := []byte(TestName)
	for i, v := range t1 {
		fmt.Println(i, string(v))
	}
	fmt.Println("=========")
	t2 := []rune(TestName2)
	for i, v := range t2 {
		fmt.Println(i, string(v))
	}
}

func UpdateStr1(s string) string {
	t1 := []byte(TestName)
	t1[2] = []byte(s)[0]
	return string(t1)
}

func UpdateStr2(s string) string {
	t2 := []rune(TestName2)
	t2[1] = []rune(s)[0]
	return string(t2)
}
