package main

import "fmt"

func main() {
	generateParenthesis(3)
}

func generateParenthesis(n int) []string {
	result := make([]string, 0)
	dfs("", 0, 0, n, &result)

	fmt.Println(result)
	return result
}

func dfs(current string, left, right int, n int, result *[]string) {
	if len(current) == n*2 {
		*result = append(*result, current)
		return
	}

	fmt.Println(current)

	if left < n {
		dfs(current+"(", left+1, right, n, result)
	}

	if left > right {
		dfs(current+")", left, right+1, n, result)
	}
}
