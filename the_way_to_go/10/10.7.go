package main

import (
	"fmt"
	"reflect"
)

type TarType struct {
	file1 bool   "this bool tag"
	file2 string "this string tag"
}

func main() {
	tt := TarType{true, "abc"}
	for i := 0; i < 2; i++ {
		refTag(tt, i)
	}
}
func refTag(tt TarType, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Printf("%v \n", ixField.Tag)

}
