package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// n5 := &TreeNode{5, nil, nil}
	// n4 := &TreeNode{4, nil, nil}
	// n6 := &TreeNode{6, nil, nil}
	// n3 := &TreeNode{3, nil, nil}
	// n7 := &TreeNode{7, nil, nil}

	// n5.Left = n4
	// n5.Right = n6
	// n6.Left = n3
	// n6.Right = n7

	// isValidBST(n5)

	n2 := &TreeNode{2, nil, nil}
	n1 := &TreeNode{1, nil, nil}
	n3 := &TreeNode{3, nil, nil}

	n2.Left = n1
	n2.Right = n3

	isValidBST(n2)
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isValidBST(root *TreeNode) bool {
	_, _, r := valid(root)
	return r
}

func valid(root *TreeNode) (int, int, bool) {
	fmt.Println("visited", root)
	if root == nil {
		return math.MinInt32 - 1, math.MaxInt32 + 1, true
	}

	leftMin, leftMax, leftValid := valid(root.Left)
	rightMin, rightMax, rightValid := valid(root.Right)

	min, max := root.Val, root.Val
	isValid := leftValid && rightValid

	if root.Left != nil {
		min = Min(min, leftMin)
		max = Max(max, leftMax)
		isValid = isValid && root.Val > leftMax
	}
	if root.Right != nil {
		min = Min(min, rightMin)
		max = Max(max, rightMax)
		isValid = isValid && root.Val < rightMin
	}

	// if leftMin != math.MinInt32-1 {
	// 	min = Min(min, leftMin)
	// }
	// if rightMin != math.MinInt32-1 {
	// 	min = Min(min, rightMin)
	// }
	// if leftMax != math.MaxInt32+1 {
	// 	max = Max(max, leftMax)
	// }
	// if rightMax != math.MaxInt32+1 {
	// 	max = Max(max, rightMax)
	// }

	// if leftMax != math.MaxInt32+1 {
	// 	isValid = isValid && root.Val > leftMax
	// }
	// if rightMin != math.MinInt32-1 {
	// 	isValid = isValid && root.Val < rightMin
	// }

	// isValid = isValid && root.Val > leftMax
	// isValid = isValid && root.Val < rightMin

	// min = Min(min, Min(leftMin, rightMin))
	// max = Max(max, Max(leftMax, rightMax))

	// fmt.Println("root.Val", root.Val, isValid, leftValid && rightValid, root.Val > leftMax, root.Val < rightMin, "|", min, max, "left:", leftMin, leftMax, "right:", rightMin, rightMax)

	return min, max, isValid
}

// func valid(root *TreeNode) bool {
// 	if root == nil {
// 		return true
// 	}

// 	isValid := true
// 	if root.Left != nil {
// 		isValid = isValid && root.Left.Val <= root.Val
// 	}
// 	if root.Right != nil {
// 		isValid = isValid && root.Right.Val >= root.Val
// 	}

// 	left, right := valid(root.Left), valid(root.Right)
// 	isValid = left && right && isValid

// 	maxLeft := findMax(root.Left)
// 	minRight := findMin(root.Right)

// 	isValid = isValid && maxLeft < root.Val
// 	isValid = isValid && minRight > root.Val

// 	return isValid
// }

// func findMax(root *TreeNode) int {
// 	if root == nil {
// 		return math.MinInt32
// 	}

// 	max := root.Val
// 	l, r := findMax(root.Left), findMax(root.Right)

// 	return Max(max, Max(l, r))
// }

// func findMin(root *TreeNode) int {
// 	if root == nil {
// 		return math.MaxInt32
// 	}

// 	min := root.Val
// 	l, r := findMin(root.Left), findMin(root.Right)

// 	return Min(min, Min(l, r))
// }

// func IsValid(root *TreeNode) (int, int, bool) {
// 	fmt.Println("visited", root)
// 	if root == nil {
// 		return math.MaxInt32, math.MinInt32, true
// 	}

// 	isValid := true
// 	if root.Left != nil {
// 		isValid = isValid && root.Left.Val <= root.Val
// 	}
// 	if root.Right != nil {
// 		isValid = isValid && root.Right.Val >= root.Val
// 	}

// 	_, leftMax, leftValid := IsValid(root.Left)
// 	rightMin, _, rightValid := IsValid(root.Right)

// 	fmt.Println(root.Val, "|", leftMax, rightMin)

// 	isValid = isValid && (leftMax < root.Val)
// 	isValid = isValid && (rightMin > root.Val)
// 	isValid = leftValid && rightValid && isValid

// 	max, min := root.Val, root.Val
// 	if leftMax != math.MaxInt32 {
// 		max, min = Max(max, leftMax), Min(min, rightMin)
// 	}
// 	if rightMin != math.MinInt32 {
// 		max, min = Max(max, rightMin), Min(min, leftMax)
// 	}
// 	fmt.Println(root.Val, "|", max, min)

// 	return max, min, isValid
// }
