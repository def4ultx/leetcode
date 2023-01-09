package main

import "fmt"

func main() {
	rob([]int{2, 3, 2}) // 4
}

func rob(nums []int) int {
	a, b := robDP(nums[1:]), nums[0]+robDP(nums[2:len(nums)-1])
	fmt.Println(a, b)
	return Max(a, b)
}

func robDP(nums []int) int {
	fmt.Println(nums)
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
