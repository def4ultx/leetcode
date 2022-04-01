package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 8
	medianSlidingWindow(arr, k)
}

func medianSlidingWindow(nums []int, k int) []float64 {
	sort.Ints(nums)
	fmt.Println(nums)

	isOdd := k%2 != 0

	result := make([]float64, 0)
	for i := 0; i < len(nums)-k+1; i++ {
		var n float64
		if isOdd {
			idx := (i*2 + k) / 2
			n = float64(nums[idx])
		} else {
			idx := (i*2 + k) / 2
			n = float64(nums[idx-1]+nums[idx]) / 2
		}
		result = append(result, n)
	}

	// fmt.Println("[-1 1 3 3 5 6]")
	// fmt.Println("[0 2 3 4 5.5]")
	fmt.Println(result)
	return result

}
