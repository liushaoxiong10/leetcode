package main

import "fmt"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var node = &ListNode{}
	lastnode := node
	add := 0
	for {
		sum := add
		add = 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		if sum >= 10 {
			sum = sum - 10
			add = 1
		}
		lastnode.Next = &ListNode{
			Val: sum,
		}
		lastnode = lastnode.Next

		if l1 == nil && l2 == nil && add == 0 {
			break
		}
	}
	return node.Next
}

func main() {
	l1 := CreateNewListNode([]int{9, 9})
	l2 := CreateNewListNode([]int{1})
	l3 := addTwoNumbers(l1, l2)
	for {
		if l3 == nil {
			break
		}
		fmt.Println(l3.Val)
		l3 = l3.Next
	}
}

// ListNode

// CreateNewListNode ...
func CreateNewListNode(vals []int) *ListNode {
	if len(vals) == 0 {
		return &ListNode{}
	}
	var l = &ListNode{}
	lastnode := l
	for _, v := range vals {
		lastnode = lastnode.AddNextNode(v)
	}
	return l.Next
}

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

// AddNextNode ...
func (l *ListNode) AddNextNode(val int) *ListNode {
	l.Next = &ListNode{
		Val: val,
	}
	return l.Next
}
