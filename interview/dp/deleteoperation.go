package main

func main() {

}

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
		dp[i][0] = i
	}
	for i := range dp[0] {
		dp[0][i] = i
	}

	for row := 1; row < len(word1)+1; row++ {
		for col := 1; col < len(word2)+1; col++ {
			if word1[row-1] == word2[col-1] {
				dp[row][col] = dp[row-1][col-1]
			} else {
				dp[row][col] = 1 + Min(dp[row-1][col], dp[row][col-1])
			}
		}
	}

	return dp[len(word1)][len(word2)]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
