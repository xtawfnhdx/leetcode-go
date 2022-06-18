package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func main() {
	start := time.Now()
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}
	where()
	i := 10
	fmt.Printf("i is %d\n", i)
	where()
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("time:%s \n", delta)
}
