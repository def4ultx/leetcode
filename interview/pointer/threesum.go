package main

import (
	"fmt"
	"sort"
)

func main() {
	threeSum([]int{-1, 0, 1, 2, -1, -4})
}

func threeSum(nums []int) [][]int {
	mapping := make(map[Pair]struct{})

	sort.Ints(nums)
	for i, v := range nums {
		pairs := twoSum(nums[i+1:], 0-v)
		for _, p := range pairs {
			p.number = v
			mapping[p] = struct{}{}
		}
	}

	fmt.Println(mapping)

	result := make([][]int, len(mapping))
	idx := 0
	for p := range mapping {
		result[idx] = []int{p.number, p.left, p.right}
		idx++
	}
	return result
}

type Pair struct {
	left, right int
	number      int
}

func twoSum(numbers []int, target int) []Pair {

	pairs := make([]Pair, 0)

	l, r := 0, len(numbers)-1
	for l < r {
		sum := numbers[l] + numbers[r]
		if sum == target {
			p := Pair{numbers[l], numbers[r], 0}
			pairs = append(pairs, p)

			l++
			r--
		} else if sum > target {
			r--
		} else {
			l++
		}
	}
	return pairs
}
