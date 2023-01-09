package main

import "fmt"

func main() {
	data := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	r := isValidSudoku(data)
	fmt.Println("result:", r)
}

func isValidSudoku(board [][]byte) bool {
	boardSize := 9

	boxes := make([][]map[byte]int, boardSize/3)
	cols := make([]map[byte]int, boardSize)
	rows := make([]map[byte]int, boardSize)

	for i := 0; i < boardSize; i++ {
		rows[i] = make(map[byte]int)
		cols[i] = make(map[byte]int)
	}

	for i := 0; i < 3; i++ {
		boxes[i] = make([]map[byte]int, 3)
		for j := 0; j < 3; j++ {
			boxes[i][j] = make(map[byte]int)
		}
	}

	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			val := board[i][j]
			if val != '.' {
				if val == 57 {
					fmt.Println("found", val, i, j, i/3, j%3)
				}
				boxes[i/3][j/3][val]++
				rows[i][val]++
				cols[j][val]++
			}
		}
	}

	fmt.Println("boxes", boxes)
	fmt.Println("rows", rows)
	fmt.Println("cols", cols)

	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			val := board[i][j]
			if val != '.' {
				count, ok := boxes[i/3][j/3][val]
				if ok && count > 1 {
					return false
				}
				count, ok = rows[i][val]
				if ok && count > 1 {
					fmt.Println("return false, rows", i, j, val)
					return false
				}
				count, ok = cols[j][val]
				if ok && count > 1 {
					fmt.Println("return false, cols", i, j, val)
					return false
				}
			}
		}
	}

	return true
}
