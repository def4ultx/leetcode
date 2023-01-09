package main

import "fmt"

func main() {
	maxCoins([]int{3, 1, 5, 8})
}

func maxCoins(nums []int) int {

	size := len(nums)

	dp := make([][]int, size)
	for i := range dp {
		dp[i] = make([]int, size)
	}
	for i := 1; i < size-1; i++ {
		dp[i][i] = nums[i-1] * nums[i] * nums[i+1]
	}
	dp[0][0] = nums[0] * nums[1]
	dp[size-1][size-1] = nums[size-1] * nums[size-2]

	fmt.Println(dp)

	for i := size - 2; i >= 0; i-- {
		for j := i + 1; j < size; j++ {
			fmt.Println(i, j)
			left, moreleft := 1, 1
			if j >= 2 {
				moreleft = nums[j-2]
			}
			if j >= 1 {
				left = nums[j-1]
			}

			a := dp[i+1][j] + left*moreleft + dp[i+1][j-1]
			b := dp[i][j-1] + nums[j]*moreleft + dp[i+1][j-1]

			fmt.Println(a, "|", dp[i+1][j], left, moreleft, nums[j], dp[i+1][j-1])
			fmt.Println(b, "|", dp[i][j-1], left, moreleft, nums[j], dp[i+1][j-1])
			fmt.Println("----")
			dp[i][j] = Max(a, b)
		}
	}

	fmt.Println(dp)

	return 1
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
