package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var te *ListNode
	var lre *ListNode

	temp, sum := (l1.Val+l2.Val)/10, (l1.Val+l2.Val)%10
	lre = &ListNode{sum, nil}
	l1, l2 = l1.Next, l2.Next

	te = lre

	for {
		var v1, v2 = 0, 0
		if l1 == nil && l2 == nil {
			break
		}
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		sum = (v1 + v2 + temp) % 10
		temp = (v1 + v2 + temp) / 10
		te.Next = &ListNode{sum, nil}
		te = te.Next
	}

	if temp == 1 {
		te.Next = &ListNode{1, nil}
	}

	return lre
}

func main() {
	var t3 = ListNode{5, nil}
	var t2 = ListNode{4, &t3}
	var t1 = ListNode{2, &t2}

	var n3 = ListNode{5, nil}
	var n2 = ListNode{6, &n3}
	var n1 = ListNode{5, &n2}

	var ss = addTwoNumbers(&t1, &n1)
	println(ss)
}
