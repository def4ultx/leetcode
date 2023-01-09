package main

func main() {

}

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for i := range dp[0] {
		dp[0][i] = 1
	}

	for row := 1; row < m; row++ {
		for col := 1; col < n; col++ {
			dp[row][col] = dp[row][col-1] + dp[col][row-1]
		}
	}

	return dp[m-1][n-1]
}
