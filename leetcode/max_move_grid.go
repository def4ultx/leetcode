package main

import "fmt"

func main() {
	maxMoves([][]int{{2, 4, 3, 5}, {5, 4, 9, 3}, {3, 4, 2, 11}, {10, 9, 13, 15}})
}

func maxMoves(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var max int
	for i := 0; i < m; i++ {
		m := dfs(grid, grid[i][0], i, 0, dp)
		max = Max(max, m)
	}
	fmt.Println("result", max)

	return max
}

func dfs(grid [][]int, prev, row, col int, dp [][]int) int {
	m, n := len(grid), len(grid[0])

	fmt.Println("visit", row, col, prev)

	if row < 0 || col < 0 || row >= m || col >= n {
		return 0
	}
	if prev >= grid[row][col] && col != 0 {
		return 0
	}
	if dp[row][col] != -1 {
		return dp[row][col]
	}

	a := dfs(grid, grid[row][col], row-1, col+1, dp)
	b := dfs(grid, grid[row][col], row, col+1, dp)
	c := dfs(grid, grid[row][col], row+1, col+1, dp)

	max := Max(a, Max(b, c)) + 1

	dp[row][col] = max
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
