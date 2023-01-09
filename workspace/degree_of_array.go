package main

import "fmt"

func main() {
	findShortestSubArray([]int{1, 2, 2, 3, 1})
	findShortestSubArray([]int{1, 2, 2, 3, 1, 4, 2})
}

func findShortestSubArray(nums []int) int {
	mapping := make(map[int]int)
	for _, v := range nums {
		mapping[v]++
	}

	var maxCount int
	var maxNum int
	for k, v := range mapping {
		if maxCount < v {
			maxCount = v
			maxNum = k
		}
	}

	l, r := -1, -1
	for i, v := range nums {
		if v == maxNum {
			r = i
			if l == -1 {
				l = i
			}
		}
	}

	fmt.Println(r, l)

	fmt.Println(r - l)

	return r - l + 1

}
