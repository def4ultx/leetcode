package main

import "fmt"

func main() {
	grid := [][]byte{
		// {1, 1, 0, 0, 0},
		// {1, 1, 0, 0, 0},
		// {0, 0, 1, 0, 0},
		// {0, 0, 0, 1, 1},
		// {0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		// {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		// {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		// {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		// {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		// {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		// {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		// {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	numIslands(grid)
}

// BFS => using queue
func numIslands(grid [][]byte) int {
	type Item struct {
		i, j int
	}
	var count int

	queue := make([]Item, 0)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 || grid[i][j] == 2 {
				continue
			}
			count++

			queue = append(queue, Item{i, j})
			for len(queue) > 0 {
				point := queue[0]
				queue = queue[1:]

				fmt.Println(point)

				if point.i < 0 || point.j < 0 || point.i >= len(grid) || point.j >= len(grid[i]) {
					continue
				}
				if grid[point.i][point.j] == 0 {
					continue
				}
				if grid[point.i][point.j] == 2 {
					continue
				}
				grid[point.i][point.j] = 2

				directionX := []int{-1, 0, 0, 1}
				directionY := []int{0, -1, 1, 0}

				for idx := 0; idx < 4; idx++ {
					item := Item{point.i + directionX[idx], point.j + directionY[idx]}
					queue = append(queue, item)
				}
			}
		}
	}
	fmt.Println(grid)
	fmt.Println(count)
	return count
}

// DFS
func numIslandsDFS(grid [][]byte) int {
	// r := make([]int, 0)
	var count int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				dfs(i, j, grid)
				count++
			}
		}
	}

	fmt.Println(grid)
	fmt.Println("result", count)
	return count
}

func dfs(i, j int, grid [][]byte) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return
	}
	if grid[i][j] == 2 {
		return
	}
	if grid[i][j] == 0 {
		return
	}

	x := []int{-1, 0, 0, 1}
	y := []int{0, 1, -1, 0}

	grid[i][j] = 2
	for idx := 0; idx < len(x); idx++ {
		dfs(i+x[idx], j+y[idx], grid)
	}
}
