package main

import (
	"fmt"
	"math"
)

func main() {
	coinChange([]int{1, 2, 5}, 11)
	// coinChange([]int{1}, 0)
}

// Need to do BFS instead

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	mapping := make(map[int]struct{})
	for _, v := range coins {
		mapping[v] = struct{}{}
	}

	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	for key := range mapping {
		if key > amount {
			continue
		}
		dp[key] = 1
	}
	for i := 0; i < len(dp); i++ {
		fmt.Println("index", i)
		for coin := range mapping {
			fmt.Println(i, coin, i-coin)
			if i-coin < 0 {
				continue
			}
			m := dp[i-coin] + 1
			dp[i] = Min(dp[i], m)
		}
	}
	fmt.Println(dp)
	fmt.Println(dp[amount])
	if dp[amount] == math.MaxInt64 {
		return -1
	}
	return dp[amount]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
