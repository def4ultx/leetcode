package main

import "fmt"

func main() {
	// lengthOfLIS([]int{1, 2, 3, 4})
	// lengthOfLIS([]int{1, 2, 4, 3})
	// lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})
	lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18, 0})
	// lengthOfLIS([]int{0, 1, 0, 3, 2, 3})
	// lengthOfLIS([]int{2, 2, 2, 2, 2})
	// lengthOfLIS([]int{4, 10, 4, 3, 8, 9})
}

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	ans := 1
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = Max(dp[i], dp[j]+1)
				ans = Max(ans, dp[i])
			}
		}
		fmt.Println(dp, nums[i])
	}
	return ans
}

func lengthOfLISDFS(nums []int) int {
	// dp := make([][]int, len(nums))
	// for i := range dp {
	// 	dp[i] = make([]int, len(nums))
	// }
	// for i := range dp {
	// 	for j := range dp[i] {
	// 		dp[i][j] = -1
	// 	}
	// }

	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = -1
	}
	l := solve(nums, -1, 0, dp)
	return l
}

func solve(nums []int, prev, idx int, dp []int) int {
	fmt.Println("visit", prev, idx, dp)
	if idx >= len(nums) {
		return 0
	}
	if dp[prev+1] != -1 {
		return dp[prev+1]
	}

	take := 0
	notTake := solve(nums, prev, idx+1, dp)
	if prev == -1 || nums[idx] > nums[prev] {
		take = 1 + solve(nums, idx, idx+1, dp)
	}

	dp[prev+1] = Max(take, notTake)

	fmt.Println(prev, idx, "value", dp[idx], take, notTake)
	return dp[prev+1]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
