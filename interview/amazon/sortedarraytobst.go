package main

func main() {

}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2

	root := &TreeNode{nums[mid], nil, nil}
	root.Left = sortedArrayToBST(nums[0:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}
