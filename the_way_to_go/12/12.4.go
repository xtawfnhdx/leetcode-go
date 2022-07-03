package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(os.Getenv("GOROOT"))
	fmt.Println(os.Getenv("GOPATH"))
	//因为工作路劲设置原因，此处的文件读取，需要写绝对路径，不然读取文件会报错
	inputFile, inputError := os.Open("/Volumes/EXPAND/Work/Git/leetcode-go/the_way_to_go/12/test.json")
	if inputError != nil {
		fmt.Println("open file error", inputError)
		return
	}
	defer inputFile.Close()

	inputReder := bufio.NewReader(inputFile)
	for {
		inputStr, readError := inputReder.ReadString('\n')
		fmt.Println(inputStr)
		if readError == io.EOF {
			return
		}
	}
}
