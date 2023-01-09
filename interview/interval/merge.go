package main

import (
	"fmt"
	"sort"
)

func main() {
	// merge([][]int{
	// 	{1, 3}, {2, 6}, {8, 10}, {15, 18},
	// })

	// merge([][]int{
	// 	{1, 4}, {4, 5},
	// })

	// merge([][]int{
	// 	{1, 4}, {5, 6},
	// })

	// merge([][]int{
	// 	{1, 4}, {5, 6}, {8, 10}, {15, 18},
	// })

	// merge([][]int{
	// 	{1, 2}, {4, 6}, {7, 12}, {8, 14}, {15, 18},
	// })

	merge([][]int{
		{2, 3}, {2, 2}, {3, 3}, {1, 3}, {5, 7}, {2, 2}, {4, 6},
	})
}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	index := 0
	for i := 1; i < len(intervals); i++ {
		prev := intervals[index]
		curr := intervals[i]

		fmt.Println(curr, prev)

		// curr.start <= prev.end
		if curr[0] <= prev[1] {
			intervals[index][1] = Max(curr[1], intervals[index][1])
			intervals[i][0] = -1 // mark as remove
		} else {
			index = i
		}
	}

	result := make([][]int, 0)
	for _, v := range intervals {
		if v[0] != -1 {
			result = append(result, v)
		}
	}

	fmt.Println("--- result")
	fmt.Println(intervals)
	fmt.Println(result)
	fmt.Println("---")

	return result
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
