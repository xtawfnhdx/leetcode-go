package main

import "fmt"

type List []int
type Appender interface {
	Append(int)
}
type Lener interface {
	Len() int
}

func (l List) Len() int {
	return len(l)
}
func (l *List) Append(val int) {
	*l = append(*l, val)
}
func CountInfo(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}
func LongEnough(l Lener) bool {
	return l.Len()*10 > 40
}
func main() {
	var lst List
	//CountInfo(lst,4,10)
	CountInfo(&lst, 5, 8)
	LongEnough(lst)
	fmt.Println(lst)

	p := new(List)
	CountInfo(p, 3, 10)
	LongEnough(p)
	fmt.Println(p)
}
