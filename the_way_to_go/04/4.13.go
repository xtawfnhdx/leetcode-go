package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "This is an example of a string"
	fmt.Printf("是否以Th开头：%t\n", strings.HasPrefix(str, "Th"))
	fmt.Printf("是否以ing结尾:%t\n", strings.HasSuffix(str, "ing"))
	fmt.Printf("是否包含an字符串:%t\n", strings.Contains(str, "an"))
	fmt.Printf("an出现的位置:%d\n", strings.Index(str, "an"))
	fmt.Printf("a出现的最后的位置:%d\n", strings.LastIndex(str, "a"))
	s1 := "this is a"
	fmt.Printf("替换后的字符串为:%s\n", strings.Replace(s1, "i", "o", -1))
	fmt.Printf("i出现过次数:%d\n", strings.Count(str, "i"))
	fmt.Printf("重复过2次数:%s\n", strings.Repeat(s1, 2))

}
