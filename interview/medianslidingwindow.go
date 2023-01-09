package main

import (
	"fmt"
)

func main() {
	// xs := []int{1, 2, 6, 9, 123, 5}
	// insertionSort(xs)
	// fmt.Println(xs)
	medianSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
}

func medianSlidingWindow(nums []int, k int) []float64 {
}

func insertionSort(arr []int) {
	if len(arr) < 1 {
		return
	}

	idx, num := len(arr)-2, arr[len(arr)-1]
	for idx >= 0 {
		fmt.Println(arr[idx], num, idx+1)
		if arr[idx] > num {
			arr[idx+1] = arr[idx]
		} else {
			break
		}
		idx--
	}
	arr[idx+1] = num
	// fmt.Println(arr, num, idx)
}
