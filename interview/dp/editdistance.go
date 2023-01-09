package main

func main() {
	minDistance("horse", "ros")
}

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word2)+1)
	for i := range dp {
		dp[i] = make([]int, len(word1)+1)
		dp[i][0] = i
	}
	for i := 0; i < len(word1)+1; i++ {
		dp[0][i] = i
	}

	for row := 1; row < len(word2)+1; row++ {
		for col := 1; col < len(word1)+1; col++ {
			if word2[row-1] == word1[col-1] {
				dp[row][col] = dp[row-1][col-1]
			} else {
				a := dp[row-1][col]
				b := dp[row][col-1]
				c := dp[row-1][col-1]

				dp[row][col] = Min(a, b, c) + 1
			}
		}
	}

	return dp[len(word2)][len(word1)]
}

func Min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	} else {
		if b < c {
			return b
		}
		return c
	}
}
