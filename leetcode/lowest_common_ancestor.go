package main

import "fmt"

func main() {

	n0 := &TreeNode{0, nil, nil}
	n1 := &TreeNode{1, nil, nil}
	n2 := &TreeNode{2, nil, nil}
	n3 := &TreeNode{3, nil, nil}
	n4 := &TreeNode{4, nil, nil}
	n5 := &TreeNode{5, nil, nil}
	n6 := &TreeNode{6, nil, nil}
	n7 := &TreeNode{7, nil, nil}
	n8 := &TreeNode{8, nil, nil}

	// [3,5,1,6,2,0,8,null,null,7,4]

	n3.Left = n5
	n3.Right = n1

	n5.Left = n6
	n5.Right = n2

	n1.Left = n0
	n1.Right = n8

	n2.Left = n7
	n2.Right = n4

	fmt.Println(lowestCommonAncestor(n3, n5, n4))

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	ok, n := exists(root, p, q)
	if !ok {
		return nil
	}
	return n
}

func exists(root, p, q *TreeNode) (bool, *TreeNode) {
	if root == nil {
		return false, nil
	}

	if root == p && root == q {
		return true, root
	}

	fmt.Println("visit", root)

	leftOk, leftNode := exists(root.Left, p, q)
	rightOk, rightNode := exists(root.Right, p, q)

	if leftNode != nil {
		return leftOk, leftNode
	}
	if rightNode != nil {
		return rightOk, rightNode
	}

	rootOk := root == p || root == q

	fmt.Println("node", root.Val, leftOk, rightOk, leftNode, rightNode, rootOk)
	if rootOk && (leftOk || rightOk) {

		fmt.Println("node", root.Val)
		return true, root
	}
	if leftOk && rightOk {
		return true, root
	}
	if rootOk || leftOk || rightOk {
		return true, nil
	}
	return false, nil
}
