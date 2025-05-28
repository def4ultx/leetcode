package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	// lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})
	// lengthOfLIS([]int{0,1,0,3,2,3})
	minimumMountainRemovals([]int{2, 1, 1, 5, 6, 2, 3, 1})
}

func minimumMountainRemovals(nums []int) int {
	forward := lengthOfLIS(nums)
	slices.Reverse(nums)
	backward := lengthOfLIS(nums)

	fmt.Println(forward)
	fmt.Println(backward)

	fmt.Println("i")
	return 0
}

func lengthOfLIS(nums []int) []int {
	var q []int
	for _, num := range nums {
		idx := sort.SearchInts(q, num)
		// fmt.Println("number", num, idx, q)
		if idx == len(q) {
			q = append(q, num)
			continue
		}
		q[idx] = num
	}
	// fmt.Println(q)

	// return len(q)
	return q
}
