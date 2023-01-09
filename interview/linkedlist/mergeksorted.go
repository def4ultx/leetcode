package main

import "fmt"

func main() {
	n1 := &ListNode{1, nil}
	n2 := &ListNode{2, nil}
	n3 := &ListNode{3, nil}
	n4 := &ListNode{4, nil}

	n5 := &ListNode{5, nil}
	n6 := &ListNode{6, nil}
	n7 := &ListNode{7, nil}
	n8 := &ListNode{8, nil}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4

	n5.Next = n6
	n6.Next = n6
	n7.Next = n8

	lists := []*ListNode{n1, n5}
	mergeKLists(lists)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	for len(lists) > 1 {
		fmt.Println("this loop")
		l := merge(lists[0], lists[1])

		lists = lists[2:]
		lists = append(lists, l)
	}

	return lists[0]
}

func merge(left, right *ListNode) *ListNode {

	head := &ListNode{}
	curr := head
	for left != nil || right != nil {
		fmt.Println("merge  loop", left, right)
		var next *ListNode
		if left == nil && right == nil {
			break // no needed, already check by first condition
		} else if left == nil {
			next = right
			right = right.Next
		} else if right == nil {
			next = left
			left = left.Next
		} else if left.Val < right.Val {
			next = left
			left = left.Next
		} else {
			next = right
			right = right.Next
		}

		curr.Next = next
		curr = next
	}
	return head.Next
}
