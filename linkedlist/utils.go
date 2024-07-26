package linkedlist

import (
	"fmt"
	"github.com/ExJuser/LeetCodeUtils/slice"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func GetLinkedListLength(head *ListNode) (length int) {
	for p := head; p != nil; p = p.Next {
		length++
	}
	return length
}

func GenerateRandomLinkedList(bound, length int) *ListNode {
	return GenerateLinkedListFromSlice(slice.GenerateIntSlice(bound, length))
}

func GenerateLinkedListFromSlice(vals []int) *ListNode {
	dummy := &ListNode{}
	p := dummy
	for _, val := range vals {
		p.Next = &ListNode{Val: val}
		p = p.Next
	}
	return dummy.Next
}

func PrintLinkedList(head *ListNode) {
	vals := make([]int, 0)
	for p := head; p != nil; p = p.Next {
		vals = append(vals, p.Val)
	}
	fmt.Println(vals)
}

func CheckLinkedListEqual(listA, listB *ListNode) bool {
	p, q := listA, listB
	for ; p != nil && q != nil; p, q = p.Next, q.Next {
		if p.Val != q.Val {
			return false
		}
	}
	return p == q
}
