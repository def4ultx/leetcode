package main

import "fmt"

func main() {
	// maxProduct([]int{2, 3, -2, 4})
	// maxProduct([]int{-2, 0, -1})
	// maxProduct([]int{3, -1, 4})
	// maxProduct([]int{-3, -1, -1})
	// maxProduct([]int{2, -5, -2, -4, 3})

	// maxProduct([]int{-2})
	maxProduct([]int{-2, 0})
}

func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	max, l, r := 0, 1, 1
	for i := 0; i < len(nums); i++ {
		fmt.Println(l, nums[i], r, nums[len(nums)-i-1], max)
		l, r = l*nums[i], r*nums[len(nums)-i-1]
		max = Max(Max(max, nums[i]), Max(l, r))
		fmt.Println("max", max)
		if l == 0 {
			l = nums[i]
		}
		if r == 0 {
			r = nums[len(nums)-i-1]
		}

	}
	return max
}

// func maxProduct(nums []int) int {
// 	if len(nums) == 1 {
// 		return nums[0]
// 	}
// 	max, min := 1, 1
// 	r := 0

// 	for _, v := range nums {
// 		if v == 0 {
// 			max, min = 1, 1
// 			continue
// 		}
// 		fmt.Println(min, max, v)

// 		max, min = Max(v, Max(max*v, min*v)), Min(v, Min(max*v, min*v))
// 		r = Max(r, Max(min, max))
// 	}
// 	fmt.Println(r)

// 	return r
// }

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func maxProduct(nums []int) int {
// 	if len(nums) == 1 {
// 		return nums[0]
// 	}
// 	maxEndHere, minEndHere, maxReset, minReset := 1, 1, 1, 1
// 	maxSoFar, minSoFar, maxResetSoFar, minResetSoFar := 0, 0, 0, 0

// 	for _, v := range nums {
// 		fmt.Println("---")
// 		// fmt.Println(maxEndHere, maxSoFar, maxEndHere*v, v)
// 		// fmt.Println(minEndHere, minSoFar, minEndHere*v, v)

// 		maxEndHere, minEndHere = maxEndHere*v, minEndHere*v
// 		maxReset = maxReset * v
// 		if maxReset > maxResetSoFar {
// 			maxResetSoFar = maxReset
// 		}
// 		if maxReset < 0 {
// 			maxReset = 1
// 		}

// 		fmt.Println(minReset, minResetSoFar, v, minReset*v)
// 		minReset = minReset * v
// 		if minReset < minResetSoFar {
// 			minResetSoFar = minReset
// 		}
// 		if minReset > 0 {
// 			minReset = v
// 		}
// 		fmt.Println(minReset, minResetSoFar, v, minReset*v)

// 		if maxEndHere > maxSoFar {
// 			maxSoFar = maxEndHere
// 		}
// 		if minEndHere < minSoFar {
// 			minSoFar = minEndHere
// 		}

// 		if maxEndHere == 0 {
// 			maxEndHere = 1
// 		}
// 		if minEndHere == 0 {
// 			minEndHere = 1
// 			minSoFar = 0
// 		}

// 	}

// 	fmt.Println("end")
// 	fmt.Println(maxSoFar, minSoFar, maxResetSoFar, minResetSoFar)

// 	return Max(maxResetSoFar, Max(maxSoFar, minSoFar))
// }
