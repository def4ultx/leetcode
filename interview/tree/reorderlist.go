package main

import "fmt"

func main() {
	n1 := &ListNode{1, nil}
	n2 := &ListNode{2, nil}
	n3 := &ListNode{3, nil}
	n4 := &ListNode{4, nil}
	n5 := &ListNode{5, nil}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5

	reorderList(n1)

	// head := n1
	// if head != nil {
	// 	fmt.Println(head.Val)
	// 	head = head.Next
	// }
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	// count
	// reverse half
	// append from head, tail until done

	length, curr := 0, head
	for curr != nil {
		length++
		curr = curr.Next
	}
	curr = head

	tail := head
	for i := 0; i < (length+1)/2; i++ {
		tail = tail.Next
	}
	tail = reverse(tail)

	curr = head
	for i := 0; i < (length-1)/2; i++ {
		fmt.Println(curr.Val, i, length/2)
		curr = curr.Next
	}
	curr.Next = nil
	fmt.Println(curr)
	curr = head

	h := head
	// fmt.Println(h, h.Next, h.Next.Next, h.Next.Next.Next)
	fmt.Println(h, h.Next, h.Next.Next)

	h = tail
	fmt.Println(h, h.Next, h.Next.Next)

	for i := 0; i < length/2; i++ {
		fmt.Println(head.Val, tail.Val)

		var t1, t2 *ListNode

		t1 = head.Next
		t2 = tail.Next

		head.Next = tail
		tail.Next = t1

		head = t1
		tail = t2
	}
	// head.Next = nil
	fmt.Println(head, tail)

	h = curr
	fmt.Println(h, h.Next, h.Next.Next, h.Next.Next.Next, h.Next.Next.Next.Next)
	// h := curr
	// for h != nil {
	// 	fmt.Println(h.Val)
	// 	h = h.Next
	// }
}

func reverse(head *ListNode) *ListNode {

	var prev *ListNode
	for {
		if head == nil {
			break
		}
		temp := head.Next
		head.Next = prev
		prev = head
		head = temp
	}
	return prev
}
