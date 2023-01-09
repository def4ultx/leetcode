package main

import "fmt"

func main() {
	// lengthOfLongestSubstring("abcabcbb")
	// lengthOfLongestSubstring("bbbbb")
	lengthOfLongestSubstring("abba")
}

func lengthOfLongestSubstring(s string) int {
	mapping := make(map[byte]int)
	max := 0

	lastIdx := 0
	for i := 0; i < len(s); i++ {
		fmt.Println("current", lastIdx, i, s[lastIdx:i+1])
		idx, ok := mapping[s[i]]
		if ok {
			lastIdx = Max(lastIdx, idx+1)
		}
		mapping[s[i]] = i
		// else, just increase the index
		max = Max(max, i-lastIdx+1)
		fmt.Println("final--", lastIdx, i, s[lastIdx:i+1], max)
	}

	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
