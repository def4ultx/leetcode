package main

import "fmt"

func main() {
	numDistinct("rabbbit", "rabbit")
}

func numDistinct(s string, t string) int {
	dp := make([][]int, len(t)+1)
	for i := range dp {
		dp[i] = make([]int, len(s)+1)
	}
	for i := range dp[0] {
		dp[0][i] = 1
	}

	for row := 1; row < len(t)+1; row++ {
		for col := 1; col < len(s)+1; col++ {
			fmt.Println(row, col, s, t)
			if s[col-1] == t[row-1] {
				dp[row][col] = dp[row-1][col-1] + dp[row][col-1]
			} else {
				dp[row][col] = dp[row][col-1]
			}
		}
	}

	fmt.Println(dp)
	return dp[len(t)][len(s)]
}
