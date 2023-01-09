package main

import "fmt"

func main() {
	wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"})
}

func wordBreak(s string, wordDict []string) []string {
	words := make(map[string]int)
	for _, w := range wordDict {
		words[w]++
	}

	// dp := make([]string, 0)

	result := make([]string, 0)
	dfs(s, "", 0, words, &result)

	fmt.Println(result)
	return result
}

func dfs(str, curr string, index int, words map[string]int, result *[]string) {

	fmt.Println("visit", curr, index, str[index:])
	if index == len(str) {
		// found solution, add it to the list
		fmt.Println("found solution", curr)
		*result = append(*result, curr[:len(curr)-1])
		return
	}
	if index > len(str) {
		// solution not found
		// is this possible?
		return
	}

	for w := range words {
		l := len(w)
		fmt.Println("try", w, str[index:], index+l, len(str))
		if index+l <= len(str) && w == str[index:index+l] {
			dfs(str, curr+w+" ", index+l, words, result)
		}
	}
	return
}
