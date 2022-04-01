package main

import (
	"fmt"
	"math"
)

func main() {
	maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
}

func maxSubArray(nums []int) int {
	maxEndHere, maxSoFar := 0, math.MinInt32

	for _, v := range nums {
		maxEndHere += v
		if maxSoFar < maxEndHere {
			maxSoFar = maxEndHere
		}
		if maxEndHere < 0 {
			maxEndHere = 0
		}
	}

	fmt.Println(maxSoFar)
	return maxSoFar
}
