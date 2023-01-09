package main

import "fmt"

func main() {
	firstMissingPositive([]int{1, 2, 0})             // 3
	firstMissingPositive([]int{1, 3, 0})             // 3
	firstMissingPositive([]int{3, 4, -1, 1})         // 2
	firstMissingPositive([]int{7, 8, 9, 10, 11, 12}) // 1
	firstMissingPositive([]int{1})                   // 2
}

func firstMissingPositive(nums []int) int {
	fmt.Println("--")
	for i, v := range nums {
		if v < 0 {
			nums[i] = 0
		}
	}
	fmt.Println(nums)

	for _, v := range nums {
		idx := Abs(v)
		if idx >= 1 && idx < len(nums)+1 {
			fmt.Println(idx, nums[idx-1])
			nums[idx-1] = -Abs(nums[idx-1])
			if nums[idx-1] == 0 {
				nums[idx-1] = -len(nums)
			}
		}
	}
	fmt.Println(nums)

	for i, v := range nums {
		if v >= 0 {
			return i + 1
		}
	}
	return len(nums)
}

func Abs(a int) int {
	if a < 0 {
		return 0 - a
	}
	return a
}
