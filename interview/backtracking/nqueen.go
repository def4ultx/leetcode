package main

import (
	"fmt"
)

func main() {
	solveNQueens(9)
}

func solveNQueens(n int) [][]string {
	board := make([]string, n)
	for i := 0; i < len(board); i++ {
		str := ""
		for idx := 0; idx < len(board); idx++ {
			str += "."
		}
		board[i] = str
	}

	// board[2][4] = 1
	// board[5][6] = 1

	// fmt.Println(isValid(board, 1, 5))
	// fmt.Println(isValid(board, 5, 3))
	r := make([][]string, 0)
	dfs(&r, board, 0)

	fmt.Println(r)

	return r
}

func dfs(result *[][]string, board []string, row int) {
	if row >= len(board) {
		// fmt.Println("count")
		// append result

		r := make([]string, len(board))
		for i := range board {
			r[i] = board[i]
		}
		*result = append(*result, r)
		return
	}

	for i := row; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if isValid(board, i, j) {
				board[i] = board[i][:j] + "Q" + board[i][j+1:]
				dfs(result, board, row+1)
				board[i] = board[i][:j] + "." + board[i][j+1:]
			}
		}
	}
}

func isValid(board []string, row, col int) bool {
	size := len(board)
	for i := 0; i < size; i++ {
		if board[row][i] == 'Q' {
			return false
		}
	}
	for i := 0; i < size; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}
	for r, c := row-1, col-1; r >= 0 && c >= 0; r, c = r-1, c-1 {
		if board[r][c] == 'Q' {
			return false
		}
	}
	for r, c := row+1, col-1; r < size && c >= 0; r, c = r+1, c-1 {
		if board[r][c] == 'Q' {
			return false
		}
	}
	for r, c := row+1, col+1; r < size && c < size; r, c = r+1, c+1 {
		if board[r][c] == 'Q' {
			return false
		}
	}
	for r, c := row-1, col+1; r >= 0 && c < size; r, c = r-1, c+1 {
		if board[r][c] == 'Q' {
			return false
		}
	}

	return true
}

// func solveNQueens(n int) [][]string {
// 	board := make([][]byte, n)
// 	for i := range board {
// 		board[i] = make([]byte, n)
// 	}
// 	for i := range board {
// 		for j := range board[i] {
// 			board[i][j] = '.'
// 		}
// 	}

// 	// board[2][4] = 1
// 	// board[5][6] = 1

// 	// fmt.Println(isValid(board, 1, 5))
// 	// fmt.Println(isValid(board, 5, 3))
// 	r := make([][]string, 0)
// 	dfs(&r, board, 0)

// 	fmt.Println(r)

// 	return r
// }

// func dfs(result *[][]string, board [][]byte, row int) {
// 	if row >= len(board) {
// 		// fmt.Println("count")
// 		// append result

// 		// r := make([]string, len(board))
// 		// for i := range board {
// 		// 	str := strings.Join(board[i], "")
// 		// 	r[i] = str
// 		// }
// 		// *result = append(*result, r)
// 		return
// 	}

// 	for i := row; i < len(board); i++ {
// 		for j := 0; j < len(board); j++ {
// 			if isValid(board, i, j) {
// 				board[i][j] = 'Q'
// 				dfs(result, board, row+1)
// 				board[i][j] = '.'
// 			}
// 		}
// 	}
// }

// func isValid(board [][]byte, row, col int) bool {
// 	size := len(board)
// 	for i := 0; i < size; i++ {
// 		if board[row][i] == 'Q' {
// 			return false
// 		}
// 	}
// 	for i := 0; i < size; i++ {
// 		if board[i][col] == 'Q' {
// 			return false
// 		}
// 	}

// 	// fmt.Println("left diag----")
// 	for r, c := row-1, col-1; r >= 0 && c >= 0; r, c = r-1, c-1 {
// 		// fmt.Println(r, c)
// 		if board[r][c] == 'Q' {
// 			return false
// 		}
// 	}
// 	for r, c := row+1, col-1; r < size && c >= 0; r, c = r+1, c-1 {
// 		// fmt.Println(r, c)
// 		if board[r][c] == 'Q' {
// 			return false
// 		}
// 	}

// 	// fmt.Println("right diag----")
// 	for r, c := row+1, col+1; r < size && c < size; r, c = r+1, c+1 {
// 		// fmt.Println(r, c)
// 		if board[r][c] == 'Q' {
// 			return false
// 		}
// 	}
// 	for r, c := row-1, col+1; r >= 0 && c < size; r, c = r-1, c+1 {
// 		// fmt.Println(r, c)
// 		if board[r][c] == 'Q' {
// 			return false
// 		}
// 	}

// 	// for i := 0; i < size; i++ {
// 	// 	r := row - i
// 	// 	c := i
// 	// 	if r >= size {
// 	// 		break
// 	// 	}
// 	// 	if board[r][c] == "Q" {
// 	// 		return false
// 	// 	}
// 	// }

// 	// fmt.Println("right diag----")
// 	// for i := 0; i < size; i++ {
// 	// 	r := size - i - 1
// 	// 	c := i + 1
// 	// 	if r == 0 {
// 	// 		break
// 	// 	}
// 	// 	if board[r][c] == "Q" {
// 	// 		return false
// 	// 	}
// 	// }

// 	return true
// }
