package main

// func main() {

// }

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	profit := 0
	min, max := prices[0], prices[0]
	for _, p := range prices {
		profit = Max(profit, p-min)
		if p > max {
			max = p
		}
		if p < min {
			min, max = p, p
		}
	}

	return profit
}

// func Max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
