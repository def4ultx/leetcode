package main

import "fmt"

func main() {
	maxProfit2(1, []int{1, 2})
}

func maxProfit2(k int, prices []int) int {

	dp := make([][][]int, 2)
	dp[0] = make([][]int, k+1)
	dp[1] = make([][]int, k+1)
	for i := 0; i < k+1; i++ {
		dp[0][i] = make([]int, len(prices))
		dp[1][i] = make([]int, len(prices))

		for j := 0; j < len(prices); j++ {
			dp[0][i][j] = -1
			dp[1][i][j] = -1
		}
	}

	fmt.Println(dp[0])
	fmt.Println(dp[1])

	var dfs func(int, int, bool) int
	dfs = func(index int, remaining int, holding bool) int {
		if index >= len(prices) {
			return 0
		}
		if remaining <= 0 {
			return 0
		}

		var dpHoldingIndex int
		if holding {
			dpHoldingIndex = 1
		} else {
			dpHoldingIndex = 0
		}
		if dp[dpHoldingIndex][remaining][index] != -1 {
			return dp[dpHoldingIndex][remaining][index]
		}

		var max int
		hold := dfs(index+1, remaining, holding)
		if holding {
			sell := dfs(index+1, remaining-1, !holding) + prices[index]
			max = Max(hold, sell)
		} else {
			buy := dfs(index+1, remaining, !holding) - prices[index]
			max = Max(hold, buy)
		}

		dp[dpHoldingIndex][remaining][index] = max
		return max
	}
	return dfs(0, k, false)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	type Point struct {
		index     int
		remaining int
		holding   bool
	}

	dp := make(map[Point]int)

	var dfs func(int, int, bool) int
	dfs = func(index int, remaining int, holding bool) int {
		if index >= len(prices) {
			return 0
		}
		if remaining <= 0 {
			return 0
		}

		point := Point{index, remaining, holding}
		if max, ok := dp[point]; ok {
			return max
		}

		var max int
		hold := dfs(index+1, remaining, holding)
		if holding {
			sell := dfs(index+1, remaining-1, !holding) + prices[index]

			max = Max(hold, sell)
		} else {
			buy := dfs(index+1, remaining, !holding) - prices[index]
			max = Max(hold, buy)
		}

		dp[point] = max
		return max
	}
	return dfs(0, 2, false)
}
