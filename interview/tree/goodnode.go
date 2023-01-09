package main

import "fmt"

func main() {
	root := &TreeNode{3, nil, nil}

	node2 := &TreeNode{1, nil, nil}
	node3 := &TreeNode{4, nil, nil}
	node4 := &TreeNode{3, nil, nil}
	node5 := &TreeNode{1, nil, nil}
	node6 := &TreeNode{5, nil, nil}

	root.Left = node2
	root.Right = node3

	node2.Left = node4

	node3.Left = node5
	node3.Right = node6

	goodNodes(root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := countNode(root.Left, root.Val)
	r := countNode(root.Right, root.Val)

	fmt.Println(l, r)

	return l + r + 1
}

func countNode(root *TreeNode, max int) int {
	if root == nil {
		return 0
	}

	l := countNode(root.Left, Max(root.Val, max))
	r := countNode(root.Right, Max(root.Val, max))

	fmt.Println(root.Val, max, l, r, root.Val > max)
	total := l + r
	if root.Val >= max {
		total++
	}
	return total
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
