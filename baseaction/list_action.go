package baseaction

import "sort"

func InitList() [][]int {
	var t1 []int
	t1 = append(t1, 2)
	t1 = append(t1, 3)

	t2 := []int{}
	t2 = append(t2, 3)
	t2 = append(t2, 4)

	//设置了长度，可以直接用下标操作
	t3 := make([]int, 5)
	t3[0] = 4
	t3[2] = 6

	t4 := make([]int, 0, 4)
	t4 = append(t4, 5)
	t4 = append(t4, 6)

	//有3个值，并同时初始化
	//res := [][]int{t1, t2, t3}

	//最大有四个值，初始化的时候，没有值，但是都只有默认值nil  类型为[]int
	res := make([][]int, 5)
	res[0] = t1
	res[1] = t2
	res[2] = t3
	res[3] = t4
	return res
}

func ListDel(s string) []string {
	t1 := []string{"a", "b", "c", "d"}
	// 切片没有删除操作，只能过滤然后寻找到某个值，然后执行切片重组
	for i := range t1 {
		if t1[i] == s {
			copy(t1[i:], t1[i+1:])
			return t1[:len(t1)-1]
		}
	}
	return t1
}

func ListInsert(li []string, i int, s string) []string {
	if i >= len(li) {
		return li
	}
	//res := make([]string, 0, len(li)+1)
	//res = append(res, li[:i]...)
	//res = append(res, s)
	//res = append(res, li[i:]...)

	//append 支持链式操作  第二个append其实创建了一个临时切片
	res := append(li[:i], append([]string{s}, li[i:]...)...)
	return res
}

func ListSortByString(li []string, reverse bool) []string {
	if reverse {
		sort.Sort(sort.Reverse(sort.StringSlice(li)))
		return li
	}
	sort.Strings(li)
	return li
}

// ListSortByInt 整数切片排序
func ListSortByInt(li []int, reverse bool) []int {
	if reverse {
		sort.Slice(li, func(x, y int) bool {
			return li[x] > li[y]
		})
		return li
	}
	sort.Ints(li)
	return li
}
