package main

import "fmt"

func main() {
	grid := [][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}
	orangesRotting(grid)
}

func orangesRotting(grid [][]int) int {
	type Pair struct {
		i, j  int
		count int
	}

	queue := make([]Pair, 0)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 2 {
				grid[i][j] = 1
				queue = append(queue, Pair{i, j, 0})
			}
		}
	}

	count := 0
	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]

		if vertex.i < 0 || vertex.j < 0 || vertex.i >= len(grid) || vertex.j >= len(grid[vertex.i]) {
			continue
		}
		if grid[vertex.i][vertex.j] == 0 {
			continue
		}
		if grid[vertex.i][vertex.j] == 2 {
			continue
		}
		grid[vertex.i][vertex.j] = 0

		directionX := []int{-1, 0, 0, 1}
		directionY := []int{0, -1, 1, 0}
		for idx := 0; idx < 4; idx++ {
			pair := Pair{
				i:     vertex.i + directionX[idx],
				j:     vertex.j + directionY[idx],
				count: vertex.count + 1,
			}
			queue = append(queue, pair)
		}
		count = Max(count, vertex.count)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	fmt.Println("count", count)
	return count
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
