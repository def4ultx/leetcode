package main

import (
	"fmt"
	"math"
)

func main() {
	grid := [][]int{
		// {1, 3, 1},
		// {1, 5, 1},
		// {4, 2, 1},

		{1, 2, 3},
		{4, 5, 6},
	}
	minPathSum(grid)
}

type Coordiation struct {
	row, col int
	sum      int
}

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dist := make([][]int, m)
	for i := range dist {
		dist[i] = make([]int, n)
	}
	for i := range dist {
		for j := range dist[0] {
			dist[i][j] = math.MaxInt32
		}
	}

	c0 := Coordiation{0, 0, 0}
	queue := make([]Coordiation, 0)
	queue = append(queue, c0)

	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]

		fmt.Println("visited", c)
		if c.row < 0 || c.col < 0 || c.row >= m || c.col >= n {
			continue
		}

		sum := c.sum + grid[c.row][c.col]
		fmt.Println("sum", sum, dist[c.row][c.col], dist)
		if sum < dist[c.row][c.col] {
			dist[c.row][c.col] = sum
			c1 := Coordiation{c.row, c.col + 1, sum}
			c2 := Coordiation{c.row + 1, c.col, sum}

			queue = append(queue, c1, c2)
		}
	}

	fmt.Println(dist)
	return dist[m-1][n-1]
}
