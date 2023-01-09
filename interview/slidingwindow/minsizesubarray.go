package main

import "fmt"

func main() {
	// minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3})
	minSubArrayLen(15, []int{1, 2, 3, 4, 5})
}

func minSubArrayLen(target int, nums []int) int {
	left, right := 0, 0
	sum, result := 0, len(nums)+1
	for right < len(nums) {
		if sum < target {
			fmt.Println("move right", sum, right, nums[right])
			sum += nums[right]
			right++
		} else {
			for left <= right && sum >= target {
				fmt.Println("update result", left, right, nums[left], sum, target, Min(result, right-left))
				result = Min(result, right-left)
				sum -= nums[left]
				left++
			}
		}
	}

	for left <= right && sum >= target {
		fmt.Println("update result", left, right, nums[left], sum, target, Min(result, right-left))
		result = Min(result, right-left)
		sum -= nums[left]
		left++
	}
	if result >= len(nums)+1 {
		return 0
	}
	return result
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
