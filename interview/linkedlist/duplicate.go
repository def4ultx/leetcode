package main

import "fmt"

func main() {
	findDuplicate([]int{1, 3, 4, 2, 2})
}

func findDuplicate(nums []int) int {
	fast, slow := 0, 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if fast == slow {
			break
		}
	}

	slow2 := 0
	for {
		slow = nums[slow]
		slow2 = nums[slow2]
		if slow == slow2 {
			break
		}
	}
	return nums[slow2]
}

func findDuplicateNegavive(nums []int) int {
	for _, v := range nums {

		fmt.Println(nums, nums[Abs(v)], Abs(v))
		n := nums[Abs(v)]
		if n < 0 {
			fmt.Println(Abs(v))
			return Abs(v)
		}
		nums[Abs(v)] = -nums[Abs(v)]
	}
	return -1
}

func Abs(x int) int {
	if x < 0 {
		return 0 - x
	}
	return x
}
