package main

import "fmt"

func main() {
	node3 := &TreeNode{
		Val: 3,
	}
	node2 := &TreeNode{
		Val:   2,
		Right: node3,
	}
	root := &TreeNode{
		Val:  1,
		Left: node2,
	}
	diameterOfBinaryTree(root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	a, b := diameter(root)
	fmt.Println("result")
	fmt.Println(a, b)
	return b
}

func diameter(root *TreeNode) (depth int, max int) {
	if root == nil {
		return 0, 0
	}

	left, maxLeft := diameter(root.Left)
	right, maxRight := diameter(root.Right)

	depth = Max(left, right) + 1
	max = Max(left+right, Max(maxLeft, maxRight))

	return depth, max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
