package main

import (
	"fmt"
	"reflect"
)

type NotknowType struct {
	s1, s2, s3 string
}

func (n NotknowType) String() string {
	return n.s1 + "-" + n.s2 + "-" + n.s3
}

var secret interface{} = NotknowType{"zh", "ce", "ce"}

func main() {
	vale := reflect.ValueOf(secret)
	typ := reflect.TypeOf(secret)
	fmt.Println(typ)
	knd := vale.Kind()
	fmt.Println(knd)
	for i := 0; i < vale.NumField(); i++ {
		fmt.Printf("Field %d: %v \n", i, vale.Field(i))
	}
	result := vale.Method(0).Call(nil)
	fmt.Println(result)
	fmt.Println(vale.Method(0).Type())

}
