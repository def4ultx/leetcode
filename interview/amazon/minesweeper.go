package main

import "fmt"

func main() {
	board := [][]byte{
		{'E', 'E', 'E', 'E', 'E'}, {'E', 'E', 'M', 'E', 'E'}, {'E', 'E', 'E', 'E', 'E'}, {'E', 'E', 'E', 'E', 'E'},
	}
	// updateBoard(board, []int{2, 1})
	updateBoard(board, []int{3, 0})
}

func updateBoard(board [][]byte, click []int) [][]byte {
	// r,c == x >

	clickRow, clickCol := click[0], click[1]

	if board[clickRow][clickCol] == 'M' {
		board[clickRow][clickCol] = 'X'
		return board
	}
	reveal(board, clickRow, clickCol)
	fmt.Println(board)

	return board
}

func reveal(board [][]byte, row, col int) {
	fmt.Println("visit", row, col)
	if row < 0 || col < 0 || row >= len(board) || col >= len(board[0]) {
		return
	}
	if board[row][col] != 'E' {
		return
	}
	// fmt.Println("wtf")

	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			r, c := row+i, col+j
			if r < 0 || c < 0 || r >= len(board) || c >= len(board[0]) {
				continue
			}
			if board[r][c] == 'M' {
				// fmt.Println("equal?")
				count++
			}
		}
	}
	// fmt.Println(count)

	if count == 0 {
		// fmt.Println("change me?", board[row][col], 'B')
		board[row][col] = 'B'
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				reveal(board, row+i, col+j)
			}
		}
	} else {
		board[row][col] = byte(count) + '0'
	}
}
