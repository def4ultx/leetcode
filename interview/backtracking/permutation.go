package main

import "fmt"

func main() {
	// permute([]int{1, 2, 3})
	permute([]int{5, 4, 6, 2})
}

func permute(nums []int) [][]int {
	visited := make([]bool, len(nums))
	result := make([][]int, 0)
	backtrack(visited, nums, []int{}, &result)
	fmt.Println(result)
	return result
}

func backtrack(visited []bool, nums []int, list []int, result *[][]int) {
	if len(list) == len(nums) {
		r := make([]int, len(list))
		copy(r, list)
		*result = append(*result, r)
		return
	}

	for i, v := range nums {
		if visited[i] {
			continue
		}

		fmt.Println(i, v)

		visited[i] = true
		list = append(list, v)

		backtrack(visited, nums, list, result)

		visited[i] = false
		list = list[:len(list)-1]
	}
}
