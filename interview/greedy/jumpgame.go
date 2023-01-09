package main

import "fmt"

func main() {
	jump([]int{2, 3, 1, 1, 4})
}

func canJump(nums []int) bool {
	goal := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if i+nums[i] >= goal {
			goal = i
		}
	}
	return goal == 0
}

func jump(nums []int) int {
	res := 0
	l, r := 0, 0

	for l < len(nums)-1 {
		max := 0
		for i := l; i <= r; i++ {
			fmt.Println("??", i, i+nums[i])
			max = Max(max, i+nums[i])
		}
		fmt.Println(l, r, max)
		l = r + 1
		r = max

		res++
	}
	return res
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
