package main

import "fmt"

func main() {
	twoSum([]int{2, 7, 11, 15}, 9)
	twoSum([]int{2, 3, 4}, 6)
}

func twoSum(numbers []int, target int) []int {

	l, r := 0, len(numbers)-1
	for l < r {
		sum := numbers[l] + numbers[r]
		if sum == target {
			return []int{l + 1, r + 1}
		} else if sum > target {
			r--
		} else {
			l++
		}

	}
	return nil
}

func twoSumMap(numbers []int, target int) []int {
	mapping := make(map[int]int)
	for i, v := range numbers {
		mapping[v] = i
	}

	fmt.Println(mapping)
	for i, v := range numbers {
		n, ok := mapping[target-v]
		if ok {
			fmt.Println("target", i+1, n+1)
			return []int{i + 1, n + 1}
		}
	}
	return nil
}
