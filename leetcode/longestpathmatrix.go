package main

import "fmt"

func main() {
	xs := [][]int{
		[]int{9, 9, 4}, []int{6, 6, 8}, []int{2, 1, 1},
	}
	longestIncreasingPath(xs)
}

func longestIncreasingPath(matrix [][]int) int {
	dp := make([][]int, len(matrix))
	for i := range dp {
		dp[i] = make([]int, len(matrix[0]))
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var max int
	for i := range matrix {
		for j := range matrix[0] {
			fmt.Println("====")
			max = Max(max, dfs(matrix, dp, i, j, -1))
		}
	}
	return max
}

func dfs(matrix, dp [][]int, row, col int, prev int) int {
	if row < 0 || row >= len(matrix) {
		return 0
	}
	if col < 0 || col >= len(matrix[0]) {
		return 0
	}

	current := matrix[row][col]
	if current <= prev {
		return 0
	}
	fmt.Println("visit", row, col, "prev", prev)
	if dp[row][col] != -1 {
		return dp[row][col]
	}

	dirX := []int{-1, 0, 0, 1}
	dirY := []int{0, -1, 1, 0}

	var max int
	for i := 0; i < 4; i++ {
		max = Max(max, dfs(matrix, dp, row+dirY[i], col+dirX[i], current)+1)
	}

	fmt.Println("max", max)

	dp[row][col] = max
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
