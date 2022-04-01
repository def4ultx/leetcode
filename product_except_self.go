package main

import "fmt"

func main() {
	productExceptSelf([]int{1, 2, 3, 4})
}

func productExceptSelf(nums []int) []int {
	fmt.Println(nums)

	size := len(nums)
	result := make([]int, size)
	for i := range result {
		result[i] = 1
	}

	left, right := 1, 1
	for i := 0; i < size; i++ {
		result[i] *= left
		left = left * nums[i]
	}
	for i := size - 1; i >= 0; i-- {
		result[i] *= right
		right = right * nums[i]
	}

	// lefts, rights := make([]int, size), make([]int, size)
	// left, right := 1, 1
	// for i := 0; i < size; i++ {
	// 	lefts[i] = left
	// 	left = left * nums[i]
	// }
	// for i := size - 1; i >= 0; i-- {
	// 	rights[i] = right
	// 	right = right * nums[i]
	// }

	// fmt.Println(lefts)
	// fmt.Println(rights)

	// for i := 0; i < size; i++ {
	// 	nums[i] = lefts[i] * rights[i]
	// }

	fmt.Println(result)

	return result
}
