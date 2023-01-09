package main

import "fmt"

func main() {

	board := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}
	gameOfLife(board)
}

func gameOfLife(board [][]int) {
	m, n := len(board), len(board[0])

	temp := make([][]int, m)
	for i := range temp {
		temp[i] = make([]int, n)
	}

	for i := range board {
		for j := range board[i] {
			temp[i][j] = board[i][j]
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			board[i][j] = mark(temp, i, j)
		}
	}

	fmt.Println(board)
}

func mark(board [][]int, row, col int) int {
	current := board[row][col]

	live := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			r, c := row+i, col+j
			if r < 0 || c < 0 || r >= len(board) || c >= len(board[0]) {
				continue
			}
			if i == 0 && j == 0 {
				continue
			}

			if board[r][c] == 1 {
				live++
			}
		}
	}

	fmt.Println("visit", row, col, "with value", current, live)

	if current == 1 && live < 2 {
		return 0
	}
	if current == 1 && live >= 2 && live <= 3 {
		return 1
	}
	if current == 1 && live > 3 {
		return 0
	}
	if current == 0 && live == 3 {
		return 1
	}
	return current
}
