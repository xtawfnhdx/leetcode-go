package main

import (
	"errors"
	"fmt"
)

var errNotFount error = errors.New("ab")

func main() {
	fmt.Printf("error:%v", errNotFount)
}
