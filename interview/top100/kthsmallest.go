package main

import "fmt"

func main() {

	n5 := &TreeNode{5, nil, nil}
	n4 := &TreeNode{4, nil, nil}
	n6 := &TreeNode{6, nil, nil}
	n3 := &TreeNode{3, nil, nil}
	n7 := &TreeNode{7, nil, nil}

	n5.Left = n4
	n5.Right = n6
	n6.Left = n3
	n6.Right = n7

	// n2 := &TreeNode{2, nil, nil}
	// n1 := &TreeNode{1, nil, nil}
	// n3 := &TreeNode{3, nil, nil}
	// n2.Left = n1
	// n2.Right = n3

	kthSmallest(n5, 2)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	stack := make([]*TreeNode, 0)
	// stack = append(stack, root)

	// while (true) {
	// 	while (root != null) {
	// 	  stack.push(root);
	// 	  root = root.left;
	// 	}
	// 	root = stack.pop();
	// 	if (--k == 0) return root.val;
	// 	root = root.right;
	//   }

	// for {
	// 	for root != nil {
	// 		stack = append(stack, root)
	// 		root = root.Left
	// 	}
	// 	root = stack[len(stack)-1]
	// 	stack = stack[:len(stack)-1]
	// 	k--
	// 	if k == 0 {
	// 		return root.Val
	// 	}
	// 	root = root.Right
	// }

	// inorder
	for {
		if root == nil && len(stack) == 0 {
			break
		}

		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Println(root.Val)

		root = root.Right
	}

	// preorder
	// for len(stack) > 0 {
	// 	top := stack[len(stack)-1]
	// 	stack = stack[:len(stack)-1]
	// 	fmt.Println(top.Val)

	// 	if top.Right != nil {
	// 		stack = append(stack, top.Right)
	// 	}
	// 	if top.Left != nil {
	// 		stack = append(stack, top.Left)
	// 	}
	// }

	return 0
}
