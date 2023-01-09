package main

import "fmt"

func main() {
	rob([]int{1, 2, 3, 1})    // 4
	rob([]int{2, 7, 9, 3, 1}) // 12
	rob([]int{9, 0, 0, 9})    // 18
	rob([]int{0, 9, 9, 0})    // 9
	rob([]int{1, 2})          // 2
	rob([]int{2, 1, 1, 2})    // 4
}

func rob(nums []int) int {
	if len(nums) < 3 {
		max := 0
		for _, v := range nums {
			if v > max {
				max = v
			}
		}
		return max
	}

	a, b := nums[0], Max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		a, b = b, Max(b, nums[i]+a)
	}
	return b
}

func rob2(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n <= 2 {
		return Max(nums[0], nums[1])
	}
	dp[n-1] = nums[n-1]
	dp[n-2] = Max(nums[n-1], nums[n-2])

	for i := n - 3; i >= 0; i-- {
		a := nums[i] + dp[i+2]
		b := dp[i+1]
		fmt.Println(i, a, b)

		dp[i] = Max(a, b)
	}

	fmt.Println(dp)

	return dp[0]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
