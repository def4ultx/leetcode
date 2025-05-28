package main

func main() {

}

func maxPathSum(root *TreeNode) int {
	_, max := dfs(root)
	return max
}

func dfs(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	left, leftMax := dfs(root.Left)
	right, rightMax := dfs(root.Right)

	maxEndHere := Max(root.Val+left, root.Val+right)
	maxEndHere = Max(maxEndHere, root.Val)

	maxSoFar := Max(maxEndHere, root.Val+left+right)
	if root.Left != nil {
		maxSoFar = Max(maxSoFar, leftMax)
	}
	if root.Right != nil {
		maxSoFar = Max(maxSoFar, rightMax)
	}

	return maxEndHere, maxSoFar
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
