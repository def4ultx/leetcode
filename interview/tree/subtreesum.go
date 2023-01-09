package main

func main() {

}

func findFrequentTreeSum(root *TreeNode) []int {
	mapping := make(map[int]int)
	subtreeSum(root, mapping)

	max, res := 0, make([]int, 0)
	for _, v := range mapping {
		if v > max {
			max = v
		}
	}
	if max == 0 {
		return res
	}

	for k, v := range mapping {
		if v == max {
			res = append(res, k)
		}
	}
	return res
}

func subtreeSum(root *TreeNode, mapping map[int]int) int {
	if root == nil {
		return 0
	}

	left, right := subtreeSum(root.Left, mapping), subtreeSum(root.Right, mapping)
	total := left + right + root.Val
	mapping[total]++
	return total
}
