package main

import (
	"fmt"
	"sort"
)

func main() {
	combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	arr := make([]int, 0)
	set := make(map[int]int)
	result := make([][]int, 0)
	combination2(candidates, arr, set, target, 0, &result)

	fmt.Println()
	fmt.Println("result")
	fmt.Println(result)
	return result
}

func combination2(candidates, arr []int, notPick map[int]int, target, current int, result *[][]int) {
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

	set := make(map[int]int)
	fmt.Println("visited", candidates, arr, notPick)
	for i, v := range candidates {

		if set[v] > 0 {
			continue
		}

		set[v]++
		// if notPick[v] > 0 {
		// 	continue
		// }

		// for idx := i; idx >= 0; idx-- {
		// 	notPick[candidates[idx]]++
		// }

		arr = append(arr, v)
		combination2(candidates[i+1:], arr, notPick, target, current+v, result)
		arr = arr[:len(arr)-1]

		// for idx := i; idx >= 0; idx-- {
		// 	notPick[candidates[idx]]--
		// }

		// combination(candidates[i+1:], arr, target, current, result)
	}
}
