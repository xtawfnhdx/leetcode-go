package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		fmt.Printf("基本循环%d \n", i)
	}
	for i, j := 1, 6; i < j; i, j = i+1, j-1 {
		fmt.Printf("多值循环 %d %d \n", i, j)
	}
	fi := 0
	for fi < 3 {
		fmt.Printf("fi is %d \n", fi)
		fi++
	}
	fi2 := 0
	for {
		if fi2 > 3 {
			break
		}
		fmt.Printf("无限循环%d \n", fi2)
		fi2++
	}
	str1 := "hello"
	for i, r := range str1 {
		fmt.Printf("字符串循环 %d,%c \n", i, r)
	}
}
