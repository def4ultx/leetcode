package main

import "fmt"

func main() {
	// longestPalindrome("babad")
	// longestPalindrome("cabbad")
	// longestPalindrome("abbc")
	longestPalindrome("abba")
}

func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}

	for i := range dp {
		dp[i][i] = true
	}
	for i := 0; i < len(s)-1; i++ {
		dp[i][i+1] = s[i] == s[i+1]
	}
	// fmt.Println(dp)
	for i := range dp {
		fmt.Println(dp[i])
	}
	fmt.Println("---")

	// for i := 0; i < len(s)-1; i++ {
	// 	for j := i + 1; j < len(s); j++ {
	// 		fmt.Println("visit", i, j)
	// 		dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
	// 	}
	// }
	for i := len(s) - 3; i >= 0; i-- {
		for j := i + 2; j < len(s); j++ {
			fmt.Println("visit", i, j)
			dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
		}
	}

	for i := range dp {
		fmt.Println(dp[i])
	}

	return ""
}

// working now
// Manacher's Algorithm o(n) - further reading
func longestPalindrome(s string) string {
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := expandCenter(s, i, i)
		len2 := expandCenter(s, i, i+1)
		l := Max(len1, len2)

		fmt.Println(len1, len2, "|", end-start+1)
		if l > end-start+1 {
			start = i - (l-1)/2
			end = i + l/2
			fmt.Println("change idx to", start, end)
		}
	}
	fmt.Println(s[start : end+1])
	return s[start : end+1]

	// l1, r1 := expandCenter(s, i, i)
	// l2, r2 := expandCenter(s, i, i+1)

	// length := (r1 - 1) - (l1 + 1) + 1
	// fmt.Println(l1, r1, "|", length, end-start+1., length > end-start+1)
	// if length > end-start+1 {
	// 	start, end = l1+1, r1-1
	// 	fmt.Println("change position", start, end, s[start:end+1])
	// }
	// fmt.Println(l2, r2)
	// if l > end-start {
	// 	fmt.Println("found", i, l)
	// 	start = i - (l-1)/2
	// 	end = i + (l+1)/2
	// }
	// }
	// fmt.Println(s, start, end, s[start:end+1])
	// return s[start:end]
	// return ""
}

func expandCenter(s string, l, r int) int {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return r - l - 1
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
