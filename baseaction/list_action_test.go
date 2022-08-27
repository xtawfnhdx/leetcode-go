package baseaction

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitList(t *testing.T) {
	res := InitList()
	for i, v := range res {
		fmt.Println(i, v)
	}
}

func TestListDel(t *testing.T) {
	res := ListDel("c")
	assert.Equal(t, []string{"a", "b", "d"}, res)
}

func TestListInsert(t *testing.T) {
	li := []string{"a", "b", "c", "d"}
	re1 := ListInsert(li, 7, "f")
	fmt.Println(re1)
	re2 := ListInsert(li, 2, "x")
	fmt.Println(re2)
}

func TestListSort(t *testing.T) {
	li := []string{"c", "a", "b", "x", "g"}
	fmt.Println(ListSortByString(li, false))
	fmt.Println(ListSortByString(li, true))

	li2 := []int{3, 6, 1, 6, 9}
	fmt.Println(ListSortByInt(li2, false))
	fmt.Println(ListSortByInt(li2, true))

}
