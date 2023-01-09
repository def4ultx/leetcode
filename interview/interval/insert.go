package main

import "fmt"

func main() {

	// {4, 8}
	// {1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16},

	// {6, 8}
	// {1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16},

	// insert([][]int{
	// 	{1, 3},
	// 	{6, 9},
	// }, []int{2, 5})

	insert([][]int{
		{1, 5},
	}, []int{5, 7})

	// insert([][]int{
	// 	{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16},
	// }, []int{4, 8})

	// insert([][]int{
	// 	{1, 2}, {12, 16},
	// }, []int{4, 8})

	// insert([][]int{
	// 	{5, 6}, {12, 16},
	// }, []int{1, 2})

	// insert([][]int{
	// 	{5, 8}, {12, 16},
	// }, []int{1, 6})

	// insert([][]int{
	// 	{5, 6}, {12, 16},
	// }, []int{18, 20})

	// insert([][]int{
	// 	{5, 6}, {12, 16},
	// }, []int{15, 20})
}

func insert(intervals [][]int, newInterval []int) [][]int {

	result := make([][]int, 0)
	for i, v := range intervals {
		if newInterval[1] < v[0] {
			// new.end < curr.start
			// append new interval and return result
			result = append(result, newInterval)
			result = append(result, intervals[i:]...)
			// fmt.Println("return", result, newInterval, intervals[i:])
			fmt.Println("return", result)
			return nil
		}
		if newInterval[0] > v[1] {
			fmt.Println("append", v)
			// new.start >= curr.start
			result = append(result, v)
		} else {
			a, b := Min(newInterval[0], v[0]), Max(newInterval[1], v[1])
			fmt.Println("set new interval", a, b)
			newInterval = []int{a, b}
		}
	}

	result = append(result, newInterval)
	fmt.Println(result)
	return result
	// idx := 0
	// for ; idx < len(intervals); idx++ {
	// 	// prev := intervals[i-1]
	// 	curr := intervals[idx]

	// 	// found overlapping
	// 	if curr[0] < newInterval[0] && curr[1] > newInterval[0] {
	// 		fmt.Println("idx break")
	// 		// intervals[idx][1] = Max(newInterval[1], intervals[idx][1])
	// 		break
	// 	}
	// }
	// fmt.Println(idx)

	// for i := index + 1; i < len(intervals); i++ {
	// 	prev := intervals[index]
	// 	curr := intervals[i]

	// 	// curr.start <= prev.end
	// 	if curr[0] <= prev[1] {
	// 		intervals[index][1] = Max(curr[1], intervals[index][1])
	// 		intervals[i][0] = -1 // mark as remove
	// 	}
	// }

	// result := make([][]int, 0)
	// for _, v := range intervals {
	// 	if v[0] != -1 {
	// 		result = append(result, v)
	// 	}
	// }

	// fmt.Println(intervals)
	// fmt.Println(result)
	// return nil
}

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
