package main

import "fmt"

func main() {
	grid := [][]int{
		// {1, 3, 1},
		// {1, 5, 1},
		// {4, 2, 1},

		{1, 2, 3},
		{4, 5, 6},
	}
	minPathSum(grid)
}

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// dp := make([][]int, m)
	// for i := range dp {
	// 	dp[i] = make([]int, n)
	// }

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			var prev int
			// if i-1 < 0 && j-1 < 0 {
			// 	prev = 0
			// } else if i-1 < 0 {
			// 	prev = dp[i][j-1]
			// } else if j-1 < 0 {
			// 	prev = dp[i-1][j]
			// } else {
			// 	prev = Min(dp[i-1][j], dp[i][j-1])
			// }
			if i-1 >= 0 && j-1 >= 0 {
				prev = Min(grid[i-1][j], grid[i][j-1])
			} else if i-1 >= 0 {
				prev = grid[i-1][j]
			} else if j-1 >= 0 {
				prev = grid[i][j-1]
			}

			grid[i][j] = prev + grid[i][j]
		}
	}

	fmt.Println(grid)
	return grid[m-1][n-1]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
