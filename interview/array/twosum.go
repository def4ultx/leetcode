package main

import "fmt"

func main() {
	twoSum([]int{3, 2, 4}, 6)
}

type Data struct {
	index []int
}

func twoSum(nums []int, target int) []int {
	mapping := make(map[int][]int)

	for i, v := range nums {
		if _, ok := mapping[v]; ok {
			mapping[v] = append(mapping[v], i)
		} else {
			mapping[v] = []int{i}
		}
	}

	fmt.Println(mapping)
	for n, i := range mapping {
		val := target - n
		list, ok := mapping[val]
		fmt.Println(list, val, ok)
		if ok && val != n {
			return []int{i[0], list[0]}
		}
		if val == n && len(list) >= 2 {
			return []int{list[0], list[1]}
		}
	}
	return []int{}
}
