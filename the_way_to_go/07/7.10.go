package main

import "fmt"

func main() {
	seasons := []string{"Spring", "Summer", "Autumn", "Winter"}
	for i, value := range seasons {
		fmt.Printf("-%d value is %s\n", i, value)
	}
	season := "zhangsan"
	for _, value := range season {
		fmt.Printf("%s\n", value)
	}
	items := [...]int{10, 20, 30, 40, 50}
	for _, item := range items {
		item *= item
	}
	fmt.Println(items)
}
