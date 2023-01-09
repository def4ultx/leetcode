package main

import "fmt"

func main() {
	minCostClimbingStairs([]int{10, 15, 20})
}

func minCostClimbingStairs(cost []int) int {
	memory := make([]int, len(cost)+1)

	memory[0] = 0
	memory[1] = 0
	cost = append(cost, 0)
	for i := 2; i < len(cost); i++ {
		a := memory[i-1] + cost[i-1]
		b := memory[i-2] + cost[i-2]
		memory[i] = Min(a, b)
	}

	// memory[0] = cost[0]
	// memory[1] = cost[1]
	// cost = append(cost, 0)
	// for i := 2; i < len(cost); i++ {
	// 	a := memory[i-1] + cost[i]
	// 	b := memory[i-2] + cost[i]
	// 	memory[i] = Min(a, b)
	// }

	fmt.Println(memory)
	return memory[len(cost)-1]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
