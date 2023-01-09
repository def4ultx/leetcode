package main

import (
	"fmt"
	"sort"
)

func main() {
	// eraseOverlapIntervals([][]int{
	// 	{1, 2}, {2, 3}, {3, 4}, {1, 3},
	// })

	// eraseOverlapIntervals([][]int{
	// 	{1, 100}, {11, 22}, {1, 11}, {2, 12},
	// })

	eraseOverlapIntervals([][]int{
		{-52, 31}, {-73, -26}, {82, 97}, {-65, -11}, {-62, -49}, {95, 99}, {58, 95}, {-31, 49}, {66, 98}, {-63, 2}, {30, 47}, {-40, -26},
	})
}
func eraseOverlapIntervals(intervals [][]int) int {

	if len(intervals) == 1 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	fmt.Println(intervals)
	count, index := 0, 0
	for i := 1; i < len(intervals); i++ {
		prev := intervals[index]
		curr := intervals[i]

		fmt.Println(curr, prev)

		// curr.start <= prev.end
		if curr[0] < prev[1] {
			fmt.Println("remove", index, i, "--", intervals[i])
			// intervals[index][1] = Max(curr[1], intervals[index][1])
			// intervals[i][0] = -1 // mark as remove

			count++
		} else {
			index = i
		}
	}

	fmt.Println(count)

	return count
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
