package main

import (
	"fmt"
)

func main() {
	// longestConsecutive([]int{100, 4, 200, 1, 3, 2})
	// longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})
	// longestConsecutive([]int{100, 7, 200, 8, 6, 3, 2})
	longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})
}

type Streak struct {
	min   int
	max   int
	count int
}

func longestConsecutive(nums []int) int {

	mapping := make(map[int]int)
	for _, v := range nums {
		mapping[v]++
	}

	memory := make(map[int]Streak)

	for k := range mapping {
		left, leftFound := memory[k-1]
		right, rightFound := memory[k+1]

		if leftFound && rightFound {
			left.max = right.max
			right.min = left.min

			total := left.count + right.count + 1
			left.count = total
			right.count = total

			memory[left.min] = left
			memory[right.max] = right
		} else if leftFound {
			left.max = k
			left.count += 1

			memory[k] = left
			memory[left.min] = left
		} else if rightFound {
			right.min = k
			right.count += 1

			memory[right.max] = right
			memory[k] = right
		} else {
			streak := Streak{
				min:   k,
				max:   k,
				count: 1,
			}
			memory[k] = streak
		}
	}

	fmt.Println(memory)
	max := 0
	for _, v := range memory {
		if v.count > max {
			max = v.count
		}
	}
	fmt.Println(max)
	return max
}
