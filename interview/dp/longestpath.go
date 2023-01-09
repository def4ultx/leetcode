package main

import "fmt"

func main() {
	matrix := [][]int{
		{9, 9, 4},
		{6, 6, 8},
		{2, 1, 1},
	}
	longestIncreasingPath(matrix)
}

func longestIncreasingPath(matrix [][]int) int {

	dp := make([][]int, len(matrix))
	for i := range dp {
		dp[i] = make([]int, len(matrix[0]))
	}

	var max int
	for i := range matrix {
		for j := range matrix[i] {
			max = Max(max, dfs(matrix, i, j, -1, dp))
		}
	}
	// dfs(matrix, 0, 0, -1, dp)
	fmt.Println(dp)
	fmt.Println(max)
	return max
}

func dfs(matrix [][]int, row, col int, prev int, dp [][]int) int {

	fmt.Println("visit", row, col)
	if row < 0 || col < 0 || row >= len(matrix) || col >= len(matrix[row]) {
		// fmt.Println("out of bound")
		return 0
	}
	if prev >= matrix[row][col] {
		// fmt.Println("return!?")
		return 0
	}
	if dp[row][col] > 0 {
		// fmt.Println("found dp")
		return dp[row][col]
	}
	fmt.Println("val", matrix[row][col], prev)

	directions := []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var max int
	for _, dir := range directions {
		// tmp := matrix[row][col]
		// matrix[row][col] = -1
		lenght := dfs(matrix, row+dir.y, col+dir.x, matrix[row][col], dp)
		// matrix[row][col] = tmp
		max = Max(max, lenght)
	}

	dp[row][col] = max + 1
	return dp[row][col]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
