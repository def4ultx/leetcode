package main

import "fmt"

func main() {
	letterCombinations("23")
	letterCombinations("")
	letterCombinations("2")
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	// mapping := map[byte][]byte{
	// 	'2': {'a', 'b', 'c'},
	// 	'3': {'d', 'e', 'f'},
	// 	'4': {'g', 'h', 'i'},
	// 	'5': {'j', 'k', 'l'},
	// 	'6': {'m', 'n', 'o'},
	// 	'7': {'p', 'q', 'r', 's'},
	// 	'8': {'t', 'u', 'v'},
	// 	'9': {'w', 'x', 'y', 'z'},
	// }
	mapping := [][]byte{
		{'a', 'b', 'c'},
		{'d', 'e', 'f'},
		{'g', 'h', 'i'},
		{'j', 'k', 'l'},
		{'m', 'n', 'o'},
		{'p', 'q', 'r', 's'},
		{'t', 'u', 'v'},
		{'w', 'x', 'y', 'z'},
	}
	result := make([]string, 0)
	dfs(mapping, digits, 0, "", &result)

	fmt.Println(result)
	return result
}

func dfs(mapping [][]byte, digits string, idx int, current string, result *[]string) {
	if idx >= len(digits) {
		*result = append(*result, current)
		return
	}

	// fmt.Println(mapping[digits[idx]])
	for _, c := range mapping[digits[idx]-'2'] {
		dfs(mapping, digits, idx+1, current+string(c), result)
	}
}
