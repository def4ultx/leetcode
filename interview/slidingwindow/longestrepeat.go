package main

import "fmt"

func main() {
	// characterReplacement("ABAB", 2)
	// characterReplacement("AABABBA", 1)
	characterReplacement("AABABBA", 2)
}

func characterReplacement(s string, k int) int {
	mapping := make([]int, 26)

	res := 0
	left, right := 0, 0
	for right < len(s) {
		mapping[s[right]-'A']++
		max := maxChar(mapping)

		fmt.Println(left, right, "|", max, mapping, (right-left+1)-max, "--", res)
		if (right-left+1)-max > k {
			mapping[s[left]-'A']--
			left++
		}
		res = Max(res, right-left+1)
		right++
		// res = Max(res, right-left+1)
		// max := maxChar(mapping)

		// if (right-left)-max <= k {
		// 	mapping[s[right]-'A']++
		// 	right++
		// } else {
		// 	mapping[s[left]-'A']--
		// 	left++
		// }
	}

	fmt.Println(res, left, right, right-left)

	return res
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxChar(list []int) int {
	max := list[0]
	for _, v := range list {
		if v > max {
			max = v
		}
	}
	return max
}
