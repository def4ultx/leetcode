package main

import "fmt"

func main() {
	// fmt.Println(0 ^ 2)
	// fmt.Println(0 ^ 3)
	// fmt.Println(0 ^ 5)

	// fmt.Println(1<<5 - 1)
	getMaximumXor([]int{0, 1, 1, 3}, 2)       // [0,3,2,3]
	getMaximumXor([]int{2, 3, 4, 7}, 3)       // [5,2,6,5]
	getMaximumXor([]int{0, 1, 2, 2, 5, 7}, 3) // [4,3,6,4,6,7]

}

func getMaximumXor(nums []int, maximumBit int) []int {

	maxK := 1<<maximumBit - 1
	// fmt.Println(maxK)

	current := nums[0]
	for i := 1; i < len(nums); i++ {
		current ^= nums[i]
	}

	results := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		fmt.Println("round", len(nums)-i-1)

		// max, kk := -1, -1
		// for k := 0; k <= maxK; k++ {

		// 	c := current ^ k
		// 	fmt.Println(current, "--", c, "from", k, "max", max)
		// 	if c > max {
		// 		max = c
		// 		kk = k
		// 	}
		// }
		// fmt.Println(maxK, current, kk)
		// if maxK > current {
		// 	results[len(nums)-i-1] = maxK - current
		// } else {
		// 	results[len(nums)-i-1] = current - maxK
		// }
		results[len(nums)-i-1] = current ^ maxK

		current ^= nums[i]
	}
	fmt.Println(results)
	return results
}
