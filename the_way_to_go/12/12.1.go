package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t1, t2 string
	fmt.Println("please inter t1,t2:")
	ar1, err1 := fmt.Scan(&t1, &t2)
	fmt.Println(t1, t2, ar1, err1)

	fmt.Println("please inter t1,t2:")
	ar3, err3 := fmt.Scanln(&t1, &t2)
	fmt.Println(t1, t2, ar3, err3)

	var x1, x2 string
	var y1 int
	ar2, err2 := fmt.Scanf("1:%s 2:%s 3:%d", &x1, &x2, &y1)
	fmt.Println(x1, x2, y1, ar2, err2)

	ar4, err4 := fmt.Sscan("zhao liu", &t1, &t2)
	fmt.Println(t1, t2, ar4, err4)

	input := "56.12 / 5212 / Go"
	format := "%f  /  %d  /  %s"
	var f float32
	var i int
	var s string
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println(f, i, s)

	var inputReader *bufio.Reader

	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("please input some:")
	input2, err5 := inputReader.ReadString('\n')
	if err5 == nil {
		fmt.Println("value is :", input2)
	}
}
