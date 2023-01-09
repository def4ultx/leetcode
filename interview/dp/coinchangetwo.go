package main

import (
	"fmt"
	"sort"
)

func main() {
	change(5, []int{1, 2, 5})
	change(11, []int{1, 2, 5})
	change(13, []int{1, 2, 5})
	change(5, []int{1, 2})
}

// Can do space: o(n) by only keeping the last row memory

func change(amount int, coins []int) int {
	if amount == 0 {
		return 1
	}
	sort.Ints(coins)

	table := make([]int, amount+1)
	table[0] = 1

	for row := 1; row <= len(coins); row++ {
		result := make([]int, amount+1)
		result[0] = 1

		for col := 0; col < amount+1; col++ {
			var include, exclude int
			if col-coins[row-1] >= 0 {
				include = result[col-coins[row-1]]
			}
			exclude = table[col]
			result[col] = include + exclude
		}
		table = result
	}
	return table[amount]
}

func changeOld(amount int, coins []int) int {
	if amount == 0 {
		return 1
	}
	sort.Ints(coins)

	table := make([][]int, len(coins)+1)
	for i := range table {
		table[i] = make([]int, amount+1)
		table[i][0] = 1
	}

	for row := 1; row <= len(coins); row++ {
		for col := 0; col < amount+1; col++ {
			var include, exclude int
			if col-coins[row-1] >= 0 {
				include = table[row][col-coins[row-1]]
			}
			exclude = table[row-1][col]
			table[row][col] = include + exclude
		}
	}
	fmt.Println(table)
	return table[len(coins)][amount]
}

func changeDP(amount int, coins []int) int {
	if amount == 0 {
		return 1
	}
	sort.Ints(coins)
	// mapping := make(map[int]struct{})
	// for _, v := range coins {
	// 	mapping[v] = struct{}{}
	// }

	// fmt.Println("data")
	// fmt.Println(mapping)
	// fmt.Println(memory)
	// fmt.Println("----")
	// memory := make(map[int]int)
	// count := searchBottomUp(coins, 0, amount, memory)

	memory := make([][]int, len(coins))
	for row := 0; row < len(coins); row++ {
		memory[row] = make([]int, amount)
		for col := 0; col < amount; col++ {
			memory[row][col] = -1
		}
	}
	count := searchDP(coins, 0, 0, amount, memory)
	fmt.Println(memory)
	fmt.Println("count", count)
	return count
}

func searchDP(coins []int, idx int, current, amount int, memory [][]int) int {
	if current == amount {
		return 1
	}
	if current > amount {
		return 0
	}

	// cache using idx and current
	if idx >= len(coins) {
		return 0
	}
	if memory[idx][current] != -1 {
		return memory[idx][current]
	}

	var count int
	for i := idx; i < len(coins); i++ {
		count += searchDP(coins, i, current+coins[i], amount, memory)
	}
	memory[idx][current] = count
	return count
}

func searchBottomUp(coins []int, current, amount int, memory map[int]int) int {
	// m, ok := memory[amount]
	fmt.Println("search amount", amount, coins)
	// if ok {
	// 	return m
	// }
	if current > amount {
		return 0
	}
	if current == amount {
		return 1
	}

	var count int
	for i, coin := range coins {
		// if amount-coin < 0 {
		// 	continue
		// }
		fmt.Println("loop", i, coin, "|", amount)
		x := searchBottomUp(coins[i:], current+coin, amount, memory)
		fmt.Println("recur", amount-coin, amount, coin, x)
		count += x
	}
	return count
}

func searchTopDown(coins []int, amount int, memory map[int]int) int {
	// m, ok := memory[amount]
	fmt.Println("search amount", amount, coins)
	// if ok {
	// 	return m
	// }
	if amount < 0 {
		return 0
	}
	if amount == 0 {
		return 1
	}

	var count int
	for i, coin := range coins {
		if amount-coin < 0 {
			continue
		}
		fmt.Println("loop", i, coin, "|", amount)
		x := searchTopDown(coins[i:], amount-coin, memory)
		fmt.Println("recur", amount-coin, amount, coin, x)
		count += x
	}
	// fmt.Println("amount", amount, "found", count)
	// memory[amount] = count
	return count
}
