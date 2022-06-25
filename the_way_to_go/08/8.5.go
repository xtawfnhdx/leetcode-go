package main

import "fmt"

func main() {
	//分配切片,容量是5
	items := make([]map[int]int, 5)
	for i := range items {
		//分配切片中每个数据的元素，长度是3
		items[i] = make(map[int]int, 3)
		//不是第三个元素，2和3只是key
		items[i][2] = 33
		items[i][3] = 34
		fmt.Printf("len is %d\n", len(items[i]))
	}
	fmt.Printf("A:value of items:%v \n", items)
	items2 := make([]map[int]int, 5)
	for _, item := range items2 {
		item = make(map[int]int, 1)
		item[1] = 5
	}
	fmt.Printf("B:value of items2:%v \n", items2)
}
