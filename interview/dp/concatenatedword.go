package main

import (
	"fmt"
	"sort"
)

func main() {

}

func findAllConcatenatedWordsInADict(words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	result := make([]string, 0)
	for i := len(words) - 1; i >= 0; i-- {
		w := words[i]
		found := wordBreak(w, words[:i])
		if found {
			result = append(result, w)
		}
	}
	return result
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
