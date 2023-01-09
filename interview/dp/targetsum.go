package main

import (
	"fmt"
	"math"
)

func main() {
	findTargetSumWays([]int{1, 1, 1, 1, 1}, 3)
}

func findTargetSumWays(nums []int, target int) int {
	total := 0
	for _, v := range nums {
		total += v
	}

	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, 2*total+1)
	}

	dp[0][nums[0]+total]++
	dp[0][-nums[0]+total]++

	fmt.Println(dp[0])
	fmt.Println("----")

	for i := 1; i < len(nums); i++ {
		fmt.Println(dp[i])
		for s := -total; s <= total; s++ {
			if dp[i-1][s+total] > 0 {
				dp[i][s+nums[i]+total] += dp[i-1][s+total]
				dp[i][s-nums[i]+total] += dp[i-1][s+total]
			}
		}
		fmt.Println(dp[i])
		fmt.Println("----")
	}

	fmt.Println(dp)

	return 0
}

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, 2*sum+1)
	}
	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = math.MinInt32
		}
	}
	return findSum(nums, 0, 0, target, dp, sum)
}

func findSum(nums []int, idx int, current, target int, dp [][]int, total int) int {
	// fmt.Println("visited", idx, "current", current, target)
	if idx >= len(nums) && current != target {
		return 0
	}
	if idx >= len(nums) && current == target {
		// fmt.Println("found", idx, current, target)
		return 1
	}
	if dp[idx][current+total] >= 0 {
		return dp[idx][current+total]
	}

	count := 0
	count += findSum(nums, idx+1, current+nums[idx], target, dp, total)
	count += findSum(nums, idx+1, current-nums[idx], target, dp, total)
	dp[idx][current+total] = count

	fmt.Println("return", count)
	return count
}

func findTargetSumWays(nums []int, target int) int {
	dp := make([]map[int]int, len(nums))
	for i := range dp {
		dp[i] = make(map[int]int)
	}
	return findSum(nums, 0, 0, target, dp)
}

func findSum(nums []int, idx int, current, target int, dp []map[int]int) int {
	// fmt.Println("visited", idx, "current", current, target)
	if idx >= len(nums) && current != target {
		return 0
	}
	if idx >= len(nums) && current == target {
		// fmt.Println("found", idx, current, target)
		return 1
	}
	if val, ok := dp[idx][current]; ok {
		return val
	}

	count := 0
	count += findSum(nums, idx+1, current+nums[idx], target, dp)
	count += findSum(nums, idx+1, current-nums[idx], target, dp)
	dp[idx][current] = count

	// fmt.Println("return", count)
	return count
}
