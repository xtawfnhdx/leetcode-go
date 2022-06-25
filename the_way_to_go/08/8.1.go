package main

import "fmt"

func main() {
	var mapList map[string]int
	var mapAssigned map[string]int
	mapList = map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]int)
	mapAssigned = mapList
	mapCreated["zhang"] = 1
	mapCreated["san"] = 2
	mapAssigned["two"] = 3
	for k, v := range mapList {
		fmt.Printf("List key is %s v is %d\n", k, v)
	}
	for k, v := range mapAssigned {
		fmt.Printf("Assis key is %s v is %d\n", k, v)
	}
	for k, v := range mapCreated {
		fmt.Printf("Cre key is %s v is %d\n", k, v)
	}
}
