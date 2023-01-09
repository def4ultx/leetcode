package main

import "fmt"

func main() {
	// // findMin([]int{4, 5, 6, 7, 0, 1, 2})
	// // findMin([]int{3, 4, 5, 1, 2})
	// // findMin([]int{11, 13, 15, 17})
	// findMin([]int{2, 1})
	// // findMin([]int{1, 2})
	// // findMin([]int{3, 1, 2})
}

func findMin(nums []int) int {
	min := nums[0]
	l, r := 0, len(nums)-1

	for l <= r {
		fmt.Println(l, r, nums[l], nums[r])
		if nums[l] <= nums[r] {
			min = Min(min, nums[l])
			break
		}

		mid := l + (r-l)/2
		min = Min(min, nums[mid])

		if nums[l] <= nums[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	fmt.Println(min, l, r)
	return min
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
