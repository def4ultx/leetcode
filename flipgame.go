package main

import (
	"math"
)

func main() {
	flipgame([]int{1, 2, 4, 4, 7}, []int{1, 3, 4, 1, 3}) // 2
	// flipgame([]int{1}, []int{1})                         // 0
	// flipgame([]int{1, 1}, []int{2, 2})                   // 1
	// flipgame([]int{1, 2}, []int{2, 1})                   // 1
}

func flipgame(fronts []int, backs []int) int {
	invalid := make(map[int]int)
	for i := 0; i < len(fronts); i++ {
		if fronts[i] == backs[i] {
			invalid[fronts[i]]++
		}
	}

	min := math.MaxInt32
	for i := 0; i < len(fronts); i++ {
		_, exist := invalid[fronts[i]]
		if fronts[i] < min && !exist {
			min = fronts[i]
		}

		_, exist = invalid[backs[i]]
		if backs[i] < min && !exist {
			min = backs[i]
		}
	}
	if min == math.MaxInt32 {
		return 0
	}
	return min
}
