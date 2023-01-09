package main

import "fmt"

func main() {
	combinationSum([]int{2, 3, 6, 7}, 7)
}

func combinationSum(candidates []int, target int) [][]int {
	arr := make([]int, 0)
	result := make([][]int, 0)
	combination(candidates, arr, target, 0, &result)

	fmt.Println("result")
	fmt.Println(result)
	return result
}

func combination(candidates, arr []int, target, current int, result *[][]int) {
	if current > target {
		return
	}
	if current == target {
		r := make([]int, len(arr))
		copy(r, arr)
		fmt.Println(arr, result)
		*result = append(*result, r)
		return
	}

	fmt.Println("visited", candidates, arr)
	for i, v := range candidates {

		arr = append(arr, v)
		combination(candidates[i:], arr, target, current+v, result)
		arr = arr[:len(arr)-1]

		// if i+1 >= len(candidates) {
		// 	continue
		// }

		// combination(candidates[i+1:], arr, target, current, result)
	}
}
