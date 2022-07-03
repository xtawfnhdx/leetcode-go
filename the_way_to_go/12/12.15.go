package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//使用接口的实际示例
	fmt.Fprintf(os.Stdout, "%s\n", "this is fmt print")
	buf := bufio.NewWriter(os.Stdout)
	fmt.Fprintf(buf, "%s\n", "this bufio print")
	buf.Flush()
}
