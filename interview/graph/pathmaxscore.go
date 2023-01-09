package main

import (
	"fmt"
	"math"
)

func main() {
	// pathsWithMaxScore([]string{"E23", "2X2", "12S"})
	// pathsWithMaxScore([]string{"EX", "XS"})
	pathsWithMaxScore([]string{"E11", "XXX", "11S"})
}

func pathsWithMaxScore(board []string) []int {
	m, n := len(board), len(board[0])
	dp := make([][]Result, m)
	for i := range dp {
		dp[i] = make([]Result, n)
	}
	for i := range dp {
		for j := range dp[i] {
			dp[i][j].Count = -1
		}
	}
	r := dfs(board, m-1, n-1, dp)

	fmt.Println(dp)
	fmt.Println(r)
	mod := int(math.Pow10(9)) + 7
	return []int{r.Sum, r.Count % mod}
}

type Result struct {
	Sum   int
	Count int
}

// E23
// 2X2
// 12S

func dfs(board []string, row, col int, dp [][]Result) Result {
	fmt.Println("visited", row, col, "-")
	if row == 0 && col == 0 {
		fmt.Println("return")
		return Result{0, 1}
	}
	if row < 0 || col < 0 {
		return Result{0, 0}
	}
	if board[row][col] == 'X' {
		return Result{0, 0}
	}
	if dp[row][col].Count != 0 {
		return dp[row][col]
	}
	val := 0
	if board[row][col] != 'S' {
		val = int(board[row][col] - '0')
	}

	temp := make([]Result, 3)
	directions := []struct{ r, c int }{{-1, 0}, {0, -1}, {-1, -1}}
	for i, d := range directions {
		temp[i] = dfs(board, row+d.r, col+d.c, dp)
		fmt.Println("finished at", row, col, "from", row+d.r, col+d.c, "with value", temp[i])
	}

	max := Max(temp[0].Sum, Max(temp[1].Sum, temp[2].Sum))
	res := Result{max, 0}
	if temp[0].Sum == res.Sum {
		res.Count += temp[0].Count
	}
	if temp[1].Sum == res.Sum {
		res.Count += temp[1].Count
	}
	if temp[2].Sum == res.Sum {
		res.Count += temp[2].Count
	}
	if res.Count != 0 {
		res.Sum += val
	}

	mod := int(math.Pow10(9)) + 7
	res.Count = res.Count % mod

	// res := Result{}
	// for _, v := range temp {
	// 	if v.Sum == 0 && v.Count == 0 {
	// 		continue
	// 	} else if v.Sum > res.Sum {
	// 		res = v
	// 	} else if v.Sum == res.Sum {
	// 		res.Count += v.Count
	// 	}
	// }
	// if res.Count != 0 {
	// 	res.Sum += val
	// }

	// const mod = 1000000000
	// res.Count = res.Count % mod

	fmt.Println("result", res, row, col)
	dp[row][col] = res
	return res
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func pathsWithMaxScore(board []string) []int {
// 	m, n := len(board), len(board[0])
// 	mapping := make(map[int]int)

// 	type Point struct {
// 		r, c int
// 		sum  int
// 	}

// 	visited := make([][]bool, m)
// 	for i := range visited {
// 		visited[i] = make([]bool, n)
// 	}

// 	queue := []Point{Point{m - 1, n - 1, 0}}
// 	for len(queue) > 0 {
// 		point := queue[0]
// 		queue = queue[1:]

// 		fmt.Println("visited", point)

// 		if point.r < 0 || point.c < 0 {
// 			continue
// 		}
// 		if board[point.r][point.c] == 'E' {
// 			mapping[point.sum]++
// 			continue
// 		}
// 		if board[point.r][point.c] == 'X' {
// 			fmt.Println("continue", point)
// 			continue
// 		}
// 		// if visited[point.r][point.c] {
// 		// 	continue
// 		// }
// 		// visited[point.r][point.c] = true

// 		if board[point.r][point.c] != 'S' {
// 			number := board[point.r][point.c] - '0'
// 			fmt.Println(number, string(board[point.r][point.c]))
// 			point.sum += int(number)
// 		}

// 		directions := []struct{ r, c int }{{-1, 0}, {0, -1}, {-1, -1}}
// 		for _, d := range directions {
// 			p := Point{
// 				r:   point.r + d.r,
// 				c:   point.c + d.c,
// 				sum: point.sum,
// 			}
// 			queue = append(queue, p)
// 		}
// 	}

// 	sum, count := 0, 0
// 	for k, v := range mapping {
// 		if k >= sum {
// 			sum, count = k, v
// 		}
// 	}
// 	fmt.Println(mapping)
// 	fmt.Println("result", sum, count)
// 	return []int{sum, count}
// }
