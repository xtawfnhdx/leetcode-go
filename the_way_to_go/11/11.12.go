package main

import (
	"fmt"
	"reflect"
)

func main() {
	var f float64 = 3.4
	v := reflect.ValueOf(f)
	fmt.Println("Canset:", v.CanSet())
	fmt.Println("type is :", v.Type())
	v = reflect.ValueOf(&f)
	fmt.Println("Canset:", v.CanSet())
	fmt.Println("type is :", v.Type())
	v = v.Elem()
	fmt.Println("Canset", v.CanSet())
	v.SetFloat(5.456)
	fmt.Println("v is :", v)
	fmt.Println("f is :", f)
}
