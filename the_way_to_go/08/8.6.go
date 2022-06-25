package main

import (
	"fmt"
	"sort"
)

var (
	barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98}
)

func main() {
	keys := make([]string, len(barVal))
	i := 0

	for k, v := range barVal {
		fmt.Printf("未排序：key=%s vaule=%d \n", k, v)
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("排序后: key=%s value=%d \n", k, barVal[k])
	}
}
