package main

import (
	"fmt"
)

func main() {
	// Union find
	longestSquareStreak([]int{4, 3, 6, 16, 8, 2})
	longestSquareStreak([]int{2, 3, 5, 6, 7})
}

// func longestSquareStreak(nums []int) int {
// 	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

// 	var (
// 		current int
// 		max     int
// 	)

// 	idx := 0
// 	for i := 0; i < len(nums); i++ {
// 		fmt.Println("visit", idx, i, nums[idx], nums[i])
// 		if nums[i] == nums[idx] {
// 			continue
// 		}

// 		if nums[idx]*nums[idx] == nums[i] {
// 			current++
// 		} else {
// 			if current > max {
// 				max = current
// 			}
// 			current = 0
// 			idx = i
// 		}
// 	}

// 	fmt.Println(max)
// 	if max == 0 {
// 		return -1
// 	}
// 	return max
// }

func longestSquareStreakDFS(nums []int) int {
	dp := make(map[int]int)
	for _, v := range nums {
		dp[v] = 0
	}

	var max int
	for _, v := range nums {
		count := dfs(v, dp)
		if count > max {
			max = count
		}
	}

	fmt.Println(dp)
	fmt.Println(max)

	if max == 1 {
		return -1
	}
	return max
}

func dfs(val int, dp map[int]int) int {
	fmt.Println("visit", val)
	next := val * val
	memo, ok := dp[next]
	if !ok {
		return 1
	}
	if memo != 0 {
		return memo + 1
	}

	count := dfs(next, dp) + 1
	fmt.Println("result of", val, count)

	dp[val] = count
	return count
}

// func longestSquareStreak(nums []int) int {
// 	dp := make(map[int]int)
// 	for _, v := range nums {
// 		dp[v] = -1
// 	}

// 	var max int
// 	for _, v := range nums {
// 		val := v

// 		count := 1
// 		for {
// 			next := val * val
// 			_, ok := dp[next]
// 			if !ok {
// 				break
// 			}
// 			val = next

// 			count++
// 		}

// 		// dp
// 		if count > max {
// 			max = count
// 		}
// 	}

// 	fmt.Println(max)

// 	if max == 1 {
// 		return -1
// 	}
// 	return max
// }
