package main

import "fmt"

func main() {
	for i := 1; i <= 25; i++ {
		GetCoordinate(i, 5)
	}
	for i := 1; i <= 36; i++ {
		GetCoordinate(i, 6)
	}

	// board := [][]int{
	// 	{-1, -1, 19, 10, -1},
	// 	{2, -1, -1, 6, -1},
	// 	{-1, 17, -1, 19, -1},
	// 	{25, -1, 20, -1, -1},
	// 	{-1, -1, -1, -1, 15},
	// }
	// snakesAndLadders(board)
}

func snakesAndLadders(board [][]int) int {
	type Node struct {
		label int
		move  int
	}

	n := len(board)
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	// visited[]

	queue := make([]Node, 1)
	queue[0] = Node{1, 0}
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]

		fmt.Println("visited top", top)

		row, col := GetCoordinate(top.label, n)
		if visited[row][col] {
			continue
		}
		if top.label == n*n {
			fmt.Println(top.move)
			return top.move
		}
		visited[row][col] = true

		for i := 1; i <= 6 && top.label+i <= n*n; i++ {
			label := top.label + i
			r, c := GetCoordinate(label, n)
			if board[r][c] != -1 {
				fmt.Println("change label to", board[r][c], r, c, label, n)
				label = board[r][c]
			}
			n := Node{
				label: label,
				move:  top.move + 1,
			}
			queue = append(queue, n)
		}

	}
	return 0
}

func GetCoordinate(label int, n int) (row int, col int) {
	row = n - 1 - (label-1)/n
	if (n%2 == 0 && row%2 == 0) || (n%2 != 0 && row%2 != 0) {
		col = n - 1 - (label-1)%n
	} else {
		col = (label - 1) % n
	}

	fmt.Println(label, "--", row, col)
	return row, col
}
