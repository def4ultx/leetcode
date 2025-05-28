package main

import "fmt"

/*
[3,2,1,5] remaining=7

[3] remaining=4
| [] remaining=7

	| [] remaining=7
	| [2] remaining=7

| [3] remaining=4

	| [3] remaining=4
		| [3] remaining=4
			| [3] remaining=4
			| [3,5] remaining=0
		| [3,1] remaining=4
			| [3,1] remaining=4
			| [3,1,5] remaining=0
	| [3,2] remaining=4
		| [3,2] remaining=4
			| [3,2] remaining=4
			| [3,2,5] remaining=0
		| [3,2,1] remaining=4
			| [3,2,1] remaining=4
			| [3,2,1,5] remaining=0
*/
func main() {
	// fmt.Println(3 | 2 | 1 | 5)
	// fmt.Println(7 ^ 3 ^ 2 ^ 1 ^ 5)
	// fmt.Println(7 & 2)

	// countMaxOrSubsetsTabular([]int{3, 1})
	// countMaxOrSubsetsTabular([]int{2, 2, 2})
	countMaxOrSubsetsTabular([]int{3, 1, 2, 5})
}

func countMaxOrSubsetsTabular(nums []int) int {

	var max int
	for _, v := range nums {
		max = max | v
	}

	dp := make([][]int, len(nums)+1)
	for i := range dp {
		dp[i] = make([]int, max+1)
		for j := range dp[i] {
			dp[i][j] = 0
		}
	}

	dp[0][0] = 1

	for i := 1; i < len(nums); i++ {
		for j := max; j >= 0; j-- {
			if i > 0 {
				dp[i][j] += dp[i-1][j]
			}
			dp[i][j] += dp[i][j|nums[i]]
		}
	}

	fmt.Println(dp)
	return dp[len(nums)][max]
}

func countMaxOrSubsets(nums []int) int {
	var max int
	for _, v := range nums {
		max = max | v
	}

	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, max+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	r := countSubset(nums, 0, 0, max, dp)
	fmt.Println(r)
	return r
}

func countSubset(nums []int, index int, current int, max int, dp [][]int) int {
	for i := 0; i < index; i++ {
		fmt.Print("  ")
	}
	fmt.Println(index, current, max)
	if index == len(nums) && current != max {
		return 0
	}
	if index == len(nums) && current == max {
		return 1
	}

	// if dp[index][current] != -1 {
	// 	return dp[index][current]
	// }

	inc := countSubset(nums, index+1, current|nums[index], max, dp)
	exc := countSubset(nums, index+1, current, max, dp)

	count := inc + exc

	dp[index][current] = count
	return count
}
