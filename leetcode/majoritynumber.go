package main

import "math"

func majorityElement(nums []int) int {
	result, count := math.MinInt, 0
	for _, v := range nums {
		if v != result {
			if count > 0 {
				count--
			} else if count == 0 {
				count++
				result = v
			}
		} else {
			count++
		}
	}
	return result
}
