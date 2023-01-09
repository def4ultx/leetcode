package main

import "fmt"

func main() {
	subsets([]int{1, 2, 3})
}

func subsets(nums []int) [][]int {
	result := make([][]int, 0)

	subset := make([]int, 0)
	var fun func(i int)
	fun = func(i int) {
		if i >= len(nums) {
			list := make([]int, len(subset))
			copy(list, subset)
			fmt.Println(list, result)
			result = append(result, list)
			return
		}

		// Add num
		subset = append(subset, nums[i])
		fun(i + 1)

		subset = subset[:len(subset)-1]
		fun(i + 1)

	}
	fun(0)
	fmt.Println(result)

	return result
}
