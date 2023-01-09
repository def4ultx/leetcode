package main

import "fmt"

func main() {
	// lengthOfLongestSubstring("abcabcbb") // 3
	// lengthOfLongestSubstring("bbbbb") // 1
	// lengthOfLongestSubstring("pwwkew") // 3
	// lengthOfLongestSubstring("au") // 2
	// lengthOfLongestSubstring("aab") // 2
	// lengthOfLongestSubstring("") // 0
	// lengthOfLongestSubstring("s") // 1
	// lengthOfLongestSubstring("dvdf") // 3
	lengthOfLongestSubstring("tmmzuxt") // 5
	// lengthOfLongestSubstring("aabaab!bb") // 3
}

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	mapping := make(map[byte]int)
	mapping[s[0]] = 0

	max, lastRepeatIdx := 1, 0
	for i := 1; i < len(s); i++ {
		idx, found := mapping[s[i]]
		if found {
			fmt.Println("change lastIdx to", i, lastRepeatIdx, idx+1)
			// max = Max(max, i-idx)
			// lastRepeatIdx = idx + 1
			lastRepeatIdx = Max(lastRepeatIdx, idx+1)
			max = Max(max, i-lastRepeatIdx+1)
		} else {
			m := Max(max, i-lastRepeatIdx+1)
			fmt.Println("update value", max, "to", m, "|", i, lastRepeatIdx)
			max = Max(max, i-lastRepeatIdx+1)
		}

		// m := Max(max, i-lastRepeatIdx+1)
		// fmt.Println("update value", max, "to", m, "|", i, lastRepeatIdx)
		// max = Max(max, i-lastRepeatIdx+1)

		mapping[s[i]] = i
	}

	fmt.Println("result", max)
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
