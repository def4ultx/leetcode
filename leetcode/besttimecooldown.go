package main

import "fmt"

func main() {
	maxProfit4([]int{1, 2, 3, 0, 2})
	fmt.Println()
	fmt.Println()
	fmt.Println()
	maxProfit5([]int{1, 2, 3, 0, 2})

}

func maxProfit4(prices []int) int {
	var (
		prevBuy  int
		prevSell int
		buy      int
		sell     int
	)

	// [buy, sell, cooldown, buy, sell]

	prevBuy = -1 * 1 << 32
	for _, price := range prices {
		fmt.Println("start loop", "price=", price, "prevBuy=", prevBuy, "prevSell=", prevSell, "buy=", buy, "sell=", sell)

		buy = max(prevSell-price, prevBuy)
		sell = max(prevBuy+price, prevSell)
		prevBuy = buy
		prevSell = sell

		fmt.Println("end loop", "price=", price, "prevBuy=", prevBuy, "prevSell=", prevSell, "buy=", buy, "sell=", sell)
	}

	// fmt.Println(sell)

	return sell
}

func maxProfit5(prices []int) int {
	var (
		prevBuy  int
		prevSell int
		buy      int
		sell     int
	)

	buy = -1 * 1 << 32
	for _, price := range prices {

		fmt.Println("start loop", "price=", price, "prevBuy=", prevBuy, "prevSell=", prevSell, "buy=", buy, "sell=", sell)

		prevBuy = buy
		buy = max(prevSell-price, prevBuy)
		prevSell = sell
		sell = max(prevBuy+price, prevSell)

		fmt.Println("end loop", "price=", price, "prevBuy=", prevBuy, "prevSell=", prevSell, "buy=", buy, "sell=", sell)
	}

	return sell
}

const (
	BuyIndexDP  = 0
	SellIndexDP = 1
)

func maxProfit3(prices []int) int {
	var previousBuy, previousSell, previousHolding int
	previousBuy = -prices[0]

	for i := 1; i < len(prices); i++ {
		buy := Max(previousBuy, previousHolding-prices[i])
		sell := previousBuy + prices[i]
		hold := Max(previousSell, previousHolding)

		previousBuy = buy
		previousSell = sell
		previousHolding = hold
	}
	return Max(previousBuy, Max(previousSell, previousHolding))
}

func maxProfit2(prices []int) int {
	n := len(prices)
	justBuy := make([]int, n)
	justSell := make([]int, n)
	holding := make([]int, n)

	justBuy[0] = -prices[0]

	for i := 1; i < n; i++ {
		justBuy[i] = Max(justBuy[i-1], holding[i-1]-prices[i])
		justSell[i] = justBuy[i-1] + prices[i]
		holding[i] = Max(justSell[i-1], holding[i-1])
	}
	return Max(justBuy[n-1], Max(justSell[n-1], holding[n-1]))
}

func maxProfit(prices []int) int {
	dp := make([][]int, 3)
	dp[BuyIndexDP] = make([]int, len(prices))
	dp[SellIndexDP] = make([]int, len(prices))

	for i := range dp[0] {
		dp[0][i] = -1
		dp[1][i] = -1
	}

	profit := dfs(prices, 0, true, dp)
	fmt.Println("profit =", profit)
	return profit
}

func dfs(prices []int, index int, canBuy bool, dp [][]int) int {
	if index >= len(prices) {
		return 0
	}

	if canBuy && dp[BuyIndexDP][index] != -1 {
		return dp[BuyIndexDP][index]
	}
	if !canBuy && dp[SellIndexDP][index] != -1 {
		return dp[SellIndexDP][index]
	}

	var buy, sell, hold int
	hold = dfs(prices, index+1, canBuy, dp)
	if canBuy {
		buy = dfs(prices, index+1, false, dp) - prices[index]
		buy = Max(buy, hold)
	} else {
		sell = dfs(prices, index+2, true, dp) + prices[index]
		sell = Max(sell, hold)
	}

	if canBuy {
		dp[BuyIndexDP][index] = buy
	} else {
		dp[SellIndexDP][index] = sell
	}

	maxProfit := Max(buy, sell)
	return maxProfit
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
