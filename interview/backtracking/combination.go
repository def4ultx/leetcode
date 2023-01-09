package main

import "fmt"

func main() {
	// [1, 3, 5]
	// [[], [1], [5], [1,5], [3], [1,3], [5,3], [1,5,3]].

	fmt.Println(combination([]int{1, 3, 5}))
}

func combination(nums []int) [][]int {
	result := make([][]int, 0)
	dfs(nums, []int{}, 0, &result)
	return result
}

func dfs(nums []int, arr []int, index int, result *[][]int) {
	if index >= len(nums) {
		temp := make([]int, len(arr))
		copy(temp, arr)
		*result = append(*result, temp)
		return
	}

	arr = append(arr, nums[index])
	dfs(nums, arr, index+1, result)
	arr = arr[:len(arr)-1]
	dfs(nums, arr, index+1, result)
}
