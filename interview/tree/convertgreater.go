package main

func main() {

}

func convertBST(root *TreeNode) *TreeNode {
	convertAndSum(root, 0)
	return root
}

func convertAndSum(root *TreeNode, rightSum int) int {
	if root == nil {
		return 0
	}

	right := convertAndSum(root.Right, rightSum)
	left := convertAndSum(root.Left, rightSum+right+root.Val)

	current := root.Val
	root.Val += right

	return left + right + current
}
