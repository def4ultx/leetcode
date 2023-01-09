package main

import "fmt"

func main() {
	// fmt.Println(removeNthFromEnd(&ListNode{1, nil}, 1))

	n1 := &ListNode{1, nil}
	n2 := &ListNode{2, nil}
	n1.Next = n2
	removeNthFromEnd(n1, 1)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	size := 0
	for curr := head; curr != nil; curr = curr.Next {
		size++
	}
	if size == n {
		return head.Next
	}

	curr, removeIdx := head, size-n
	for i := 0; i < removeIdx-1 && curr != nil; i++ {
		curr = curr.Next
	}

	prev := curr
	curr = curr.Next
	if curr != nil {
		prev.Next = curr.Next
	}

	fmt.Println(head, prev, curr)

	return head
}
