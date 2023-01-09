package main

func main() {

}

func pacificAtlantic(heights [][]int) [][]int {
	if len(heights) == 0 {
		return [][]int{}
	}

	m, n := len(heights), len(heights[0])
	pacific := make([][]bool, m)
	atlantic := make([][]bool, m)

	for i := range pacific {
		pacific[i] = make([]bool, n)
		atlantic[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		dfs(heights, -1, i, 0, pacific)
		dfs(heights, -1, i, n-1, atlantic)
	}
	for i := 0; i < n; i++ {
		dfs(heights, -1, 0, i, pacific)
		dfs(heights, -1, m-1, i, atlantic)
	}

	result := make([][]int, 0)
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if pacific[r][c] && atlantic[r][c] {
				result = append(result, []int{r, c})
			}
		}
	}
	return result
}

func dfs(heights [][]int, prev int, r, c int, dp [][]bool) {
	if r < 0 || c < 0 || r >= len(heights) || c >= len(heights[0]) {
		return
	}
	if dp[r][c] {
		return
	}
	if heights[r][c] < prev {
		return
	}
	dp[r][c] = true

	directions := []struct{ r, c int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, v := range directions {
		dfs(heights, heights[r][c], r+v.r, c+v.c, dp)
	}
}
