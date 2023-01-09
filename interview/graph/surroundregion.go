package main

import "fmt"

func main() {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	solve(board)
}

func solve(board [][]byte) {
	// dfs from border and mark as visited
	// loop to flip unvisited node
	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		dfs(board, i, 0, visited)
		dfs(board, i, n-1, visited)
	}
	for i := 0; i < n; i++ {
		dfs(board, 0, i, visited)
		dfs(board, m-1, i, visited)
	}

	fmt.Println(visited)

	for r := range visited {
		for c := range visited[0] {
			if !visited[r][c] && board[r][c] == 'O' {
				fmt.Println(r, c)
				board[r][c] = 'X'
			}
		}
	}
	fmt.Println(board)
}

func dfs(board [][]byte, row, col int, visited [][]bool) {
	if row < 0 || col < 0 || row >= len(board) || col >= len(board[0]) {
		return
	}
	if visited[row][col] {
		return
	}
	if board[row][col] == 'X' {
		return
	}
	visited[row][col] = true

	directions := []struct{ r, c int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, v := range directions {
		dfs(board, row+v.r, col+v.c, visited)
	}
}
