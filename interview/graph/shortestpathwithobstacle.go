package main

func main() {

}

func shortestPath(grid [][]int, k int) int {
	type Point struct {
		row, col int
		k        int
		step     int
	}

	m, n := len(grid), len(grid[0])
	visited := make([][][]bool, m)
	for i := range visited {
		visited[i] = make([][]bool, n)
		for j := range visited[i] {
			visited[i][j] = make([]bool, k+1)
		}
	}

	queue := make([]Point, 0)
	queue = append(queue, Point{0, 0, k, 0})

	for len(queue) > 0 {
		point := queue[0]
		queue = queue[1:]

		if point.row == m-1 && point.col == n-1 {
			return point.step
		}

		if point.row < 0 || point.col < 0 || point.row >= len(grid) || point.col >= len(grid[0]) {
			continue
		}
		if visited[point.row][point.col][point.k] {
			continue
		}
		visited[point.row][point.col][point.k] = true

		if grid[point.row][point.col] == 1 && point.k == 0 {
			continue
		}

		if grid[point.row][point.col] == 1 && point.k > 0 {
			point.k--
		}

		directions := []struct{ r, c int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		for _, v := range directions {
			p := Point{
				row:  point.row + v.r,
				col:  point.col + v.c,
				k:    point.k,
				step: point.step + 1,
			}
			queue = append(queue, p)
		}
	}

	return -1
}
