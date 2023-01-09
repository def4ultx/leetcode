package main

import "fmt"

func main() {
	xs := []int{1, 4, 3, 2, 5, 2}
	list := makelist(xs)
	partition(list, 3)
}

func makelist(xs []int) *ListNode {
	head := &ListNode{}
	curr := head
	for i := 0; i < len(xs); i++ {
		curr.Next = &ListNode{xs[i], nil}
		curr = curr.Next
	}
	fmt.Println(head, head.Next, head.Next.Next)
	return head.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	l, r := &ListNode{}, &ListNode{}
	left, right := l, r

	for head != nil {
		fmt.Println(head.Val)
		if head.Val < x {
			left.Next = head
			left = left.Next
		} else {
			right.Next = head
			right = right.Next
		}
		head = head.Next
	}
	left.Next, right.Next = r.Next, nil

	left, right = l.Next, r.Next
	fmt.Println("----")
	for left != nil {
		fmt.Println(left.Val)
		left = left.Next
	}
	// fmt.Println("----")
	// for right != nil {
	// 	fmt.Println(right.Val)
	// 	right = right.Next
	// }

	// left.Next = r.Next

	// left = l
	// for left.Next != nil {
	// 	left = left.Next
	// }
	// left.Next = r.Next

	// fmt.Println(left, left.Next, r.Next)

	// left = l
	// for left != nil {
	// 	fmt.Println(left.Val, "--")
	// 	left = left.Next
	// }
	return l.Next
}
