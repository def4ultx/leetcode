package main

import "fmt"

func main() {
	knapsack([]int{5, 3, 4, 2}, []int{60, 50, 70, 30}, 5)
}

func knapsack(weights, items []int, maximum int) int {

	dp := make([][]int, len(items)+1)
	for i := range dp {
		dp[i] = make([]int, maximum+1)
	}

	for i := 1; i < len(items)+1; i++ {
		for j := 1; j < maximum+1; j++ {
			dp[i][j] = dp[i-1][j] // nopick
			fmt.Println(j)
			if j >= weights[i-1] {
				pick := dp[i-1][j-weights[i-1]] + items[i-1]
				dp[i][j] = Max(pick, dp[i][j])
			}
		}
	}

	fmt.Println(dp)
	return 0
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
