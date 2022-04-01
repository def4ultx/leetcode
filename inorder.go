package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	arr := make([]int, 0)
	l := inorderTraversal(root.Left)
	if l != nil {
		arr = append(arr, l...)
	}

	arr = append(arr, root.Val)

	l = inorderTraversal(root.Right)
	if l != nil {
		arr = append(arr, l...)
	}

	return arr
}
