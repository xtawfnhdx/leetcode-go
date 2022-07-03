package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fName := "/Volumes/EXPAND/Work/Git/leetcode-go/the_way_to_go/12/12.7.txt"
	outputFile, outputError := os.OpenFile(fName, os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Println("open file error")
		return
	}
	defer outputFile.Close()
	outputWrite := bufio.NewWriter(outputFile)
	for i := 0; i < 10; i++ {
		outputWrite.WriteString("this is test" + string(i) + "\n")
	}
	outputWrite.Flush()
	fmt.Fprintf(outputFile, "fmt text output \n")
	outputFile.WriteString("this outputFile out put \n")

}
