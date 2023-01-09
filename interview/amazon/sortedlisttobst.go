package main

import "fmt"

func main() {
	n1 := &ListNode{-10, nil}
	n2 := &ListNode{-3, nil}
	n3 := &ListNode{0, nil}
	n4 := &ListNode{5, nil}
	n5 := &ListNode{9, nil}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5

	sortedListToBST(n1)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	size := 0
	for curr := head; curr != nil; curr = curr.Next {
		size++
	}

	curr := head
	var build func(left, right int) *TreeNode
	build = func(left, right int) *TreeNode {
		fmt.Println("build", left, right)
		if left > right {
			return nil
		}

		mid := left + (right-left)/2
		l := build(left, mid-1)

		root := &TreeNode{curr.Val, nil, nil}
		curr = curr.Next

		root.Left = l
		root.Right = build(mid+1, right)

		return root
	}

	return build(0, size-1)
}

// func sortedListToBST(head *ListNode) *TreeNode {
// 	if head == nil {
// 		return nil
// 	}

// 	temp := make([]int, 0)
// 	for curr := head; curr != nil; curr = curr.Next {
// 		temp = append(temp, curr.Val)
// 	}

// 	return sortedArrayToBST(temp)
// }

// func sortedArrayToBST(nums []int) *TreeNode {
// 	if len(nums) == 0 {
// 		return nil
// 	}
// 	mid := len(nums) / 2

// 	root := &TreeNode{nums[mid], nil, nil}
// 	root.Left = sortedArrayToBST(nums[0:mid])
// 	root.Right = sortedArrayToBST(nums[mid+1:])
// 	return root
// }
