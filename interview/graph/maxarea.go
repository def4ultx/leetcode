package main

func main() {
	grid := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}
	maxAreaOfIsland(grid)
}

func maxAreaOfIsland(grid [][]int) int {
	// r := make([]int, 0)
	var max int
	for i, g := range grid {
		for j, v := range g {
			if v == -1 {
				continue
			}
			if v == 1 {
				max = Max(max, dfs(i, j, grid))
				// xx := dfs(i, j, grid)
				// fmt.Println("i, j", i, j, xx)
				// r = append(r, xx)
			}
		}
	}

	// max = dfs(0, 2, grid)
	// fmt.Println(grid)
	// fmt.Println("---")
	// fmt.Println(r)
	// fmt.Println(max)
	return max
}

func dfs(i, j int, grid [][]int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return 0
	}
	if grid[i][j] == -1 {
		return 0
	}
	if grid[i][j] == 0 {
		return 0
	}

	// fmt.Println("dfs", i, j)

	x := []int{-1, 0, 0, 1}
	y := []int{0, 1, -1, 0}

	grid[i][j] = -1
	// fmt.Println("here")
	var count int
	for idx := 0; idx < len(x); idx++ {
		// fmt.Println("idx")
		// fmt.Println(i+x[idx], j+y[idx])
		count += dfs(i+x[idx], j+y[idx], grid)
	}
	return count + 1
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
