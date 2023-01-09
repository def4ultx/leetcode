package main

func main() {

}

type Point struct {
	row, col int
}

func wallsAndGates(rooms *[][]int) {
	// write your code here

	m, n := len(*rooms), len((*rooms)[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	queue := make([]Point, 0)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (*rooms)[i][j] == 0 {
				p := Point{i, j}
				queue = append(queue, p)
			}
		}
	}

	AddRoom := func(row, col int) {
		if row < 0 || col < 0 || row >= m || col >= n {
			return
		}
		if visited[row][col] {
			return
		}
		if (*rooms)[row][col] == 0 || (*rooms)[row][col] == -1 {
			return
		}
		p := Point{row, col}
		queue = append(queue, p)
	}

	dist := 0
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			p := queue[0]
			queue = queue[1:]

			visited[p.row][p.col] = true
			AddRoom(p.row+1, p.col)
			AddRoom(p.row-1, p.col)
			AddRoom(p.row, p.col+1)
			AddRoom(p.row, p.col-1)
		}
		dist++
	}
}
