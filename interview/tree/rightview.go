package main

import "fmt"

func main() {

	n1 := &TreeNode{1, nil, nil}
	n2 := &TreeNode{2, nil, nil}
	n3 := &TreeNode{3, nil, nil}
	n4 := &TreeNode{4, nil, nil}
	n5 := &TreeNode{5, nil, nil}

	n1.Left = n2
	n1.Right = n3
	n2.Right = n5
	n3.Right = n4

	rightSideView(n1)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	result := make([]int, 0)
	dfs(root, &result, 0)
	return result
}

func dfs(root *TreeNode, result *[]int, level int) {
	if root == nil {
		return
	}

	if len(*result) < level+1 {
		*result = append(*result, root.Val)
	}

	dfs(root.Right, result, level+1)
	dfs(root.Left, result, level+1)
}

func rightSideView(root *TreeNode) []int {
	// level order tranversal
	// with level attach in queue item, set the result[level] = value

	result := make([]int, 0)

	type Node struct {
		*TreeNode
		level int
	}
	queue := []Node{Node{root, 0}}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.TreeNode == nil {
			fmt.Println("still continue")
			continue
		}
		if len(result) < node.level+1 {
			result = append(result, node.Val)
		}
		queue = append(queue, Node{node.Right, node.level + 1})
		queue = append(queue, Node{node.Left, node.level + 1})
	}
	fmt.Println(result)
	return result
}
