package main

import "fmt"

func main() {
	map1 := make(map[string]int)
	map1["ma1"] = 22
	map1["ma2"] = 23
	map1["ma3"] = 24
	v, ok := map1["ma4"]
	if ok {
		fmt.Printf("map1 ma4 vaule is %d\n", v)
	}

	delete(map1, "ma1")
	if value, ok := map1["ma1"]; ok {
		fmt.Printf("map1 ma1 value is %d\n", value)
	} else {
		fmt.Printf("map1 not contains ma1")
	}
}
