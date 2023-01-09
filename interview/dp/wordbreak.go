package main

import "fmt"

func main() {
	// wordBreak("leetcode", []string{"leet", "code"})
	// wordBreak("leetcode", []string{"lee", "tcode", "code"})
	wordBreak("goalspecial", []string{"go", "goal", "goals", "special"})
}

func wordBreak(s string, wordDict []string) bool {
	words := make(map[string]struct{})
	for _, v := range wordDict {
		words[v] = struct{}{}
	}

	result := make([]bool, len(s))
	result = append(result, true)
	for i := len(s) - 1; i >= 0; i-- {
		for w := range words {
			fmt.Println("index", i, w, len(w))
			if len(w)+i <= len(s) && s[i:i+len(w)] == w {
				if !result[i] {
					result[i] = result[i+len(w)]
				}
			}
		}
	}
	fmt.Println(result)
	return result[0]
}
