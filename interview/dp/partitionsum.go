package main

import "fmt"

func main() {
	canPartition([]int{1, 5, 11, 5})
	// canPartition([]int{1, 2, 3, 5})
}

func canPartition(nums []int) bool {

	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}

	dp := make([][]bool, len(nums)+1)
	for i := range dp {
		dp[i] = make([]bool, (sum/2)+1)
	}
	dp[0][0] = true

	for i := 1; i < len(nums)+1; i++ {
		for j := 1; j < (sum/2)+1; j++ {
			fmt.Println("visit", i, j, "|", nums[i-1], dp[i-1][j])
			dp[i][j] = dp[i-1][j]
			if j >= nums[i-1] {
				dp[i][j] = dp[i][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}

	fmt.Println(dp)
	return dp[len(nums)][(sum / 2)]
}

func canPartition2(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}

	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, (sum/2)+1)
	}
	r := dfs(nums, 0, sum, 0, 0, dp)
	return r
}

func dfs(nums []int, index int, sum, left, right int, dp [][]int) bool {

	if left == right && index == len(nums) {
		return true
	}
	if left > sum/2 || right > sum/2 {
		return false
	}
	if index >= len(nums) {
		return false
	}
	if dp[index][left] != 0 {
		return dp[index][left] == 1
	}

	v := nums[index]
	l := dfs(nums, index+1, sum, left+v, right, dp)
	r := dfs(nums, index+1, sum, left, right+v, dp)
	found := l || r
	if found {
		dp[index][left] = 1
	} else {
		dp[index][left] = -1
	}

	return found
}

// func canPartition2(nums []int) bool {
// 	sum := 0
// 	for _, v := range nums {
// 		sum += v
// 	}
// 	if sum%2 != 0 {
// 		return false
// 	}

// 	dp := make([]map[int]int, len(nums))
// 	for i := range dp {
// 		dp[i] = make(map[int]int)
// 	}
// 	r := dfs(nums, 0, sum, 0, 0, dp)
// 	return r
// }

// func dfs(nums []int, index int, sum, left, right int, dp []map[int]int) bool {

// 	if left == right && index == len(nums) {
// 		return true
// 	}
// 	if left > sum/2 || right > sum/2 {
// 		return false
// 	}
// 	if index >= len(nums) {
// 		return false
// 	}
// 	if dp[index][left] != 0 {
// 		return dp[index][left] == 1
// 	}

// 	v := nums[index]
// 	l := dfs(nums, index+1, sum, left+v, right, dp)
// 	r := dfs(nums, index+1, sum, left, right+v, dp)
// 	found := l || r
// 	if found {
// 		dp[index][left] = 1
// 	} else {
// 		dp[index][left] = -1
// 	}

// 	return found
// }
