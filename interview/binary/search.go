package main

import "fmt"

func main() {
	// fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
	fmt.Println(search([]int{1, 2}, 1))
	fmt.Println(search([]int{1, 2}, 2))
}

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for r >= l {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}
