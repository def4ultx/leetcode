package main

import "fmt"

func main() {
	node1 := &ListNode{1, nil}
	node2 := &ListNode{2, nil}
	node3 := &ListNode{3, nil}
	node4 := &ListNode{4, nil}
	node5 := &ListNode{5, nil}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	head := reverseList(node1)
	for {
		if head == nil {
			break
		}
		fmt.Println(head.Val)
		head = head.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// current := head
// var next *ListNode
// for {
// 	if current == nil {
// 		break
// 	}

// 	t := current.Next
// 	current.Next = next
// 	next = current
// 	current = t
// 	fmt.Println(current, next)
// }
// fmt.Println("---", next.Val)
// return next

func reverseList(head *ListNode) *ListNode {

	// current := head
	var prev *ListNode
	for {
		// fmt.Println(head.Val, head.Next)

		if head == nil {
			break
		}
		temp := head.Next
		// fmt.Println("head", head)
		head.Next = prev
		// fmt.Println("head", head)

		prev = head
		head = temp

		// fmt.Println(head.Val, head.Next.Val)
	}
	return prev
}
