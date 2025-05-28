package main

import "container/heap"

func main() {

}

func minimumObstacles(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	// bfs with pq??
	xs := &PointHeap{}
	heap.Init(xs)
	heap.Push(xs, Point{0, 0, grid[0][0]})

	dirX := []int{-1, 0, 0, 1}
	dirY := []int{0, -1, 1, 0}

	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[i]))
	}

	m, n := len(grid)-1, len(grid[0])-1
	for {
		pop := heap.Pop(xs).(Point)
		if visited[pop.row][pop.col] {
			continue
		}
		visited[pop.row][pop.col] = true

		k := pop.k + grid[pop.row][pop.col]
		if pop.row == m && pop.col == n {
			return k
		}

		for i := 0; i < 4; i++ {
			r := pop.row + dirY[i]
			c := pop.col + dirX[i]
			if r < 0 || r > m || c < 0 || c > n {
				continue
			}

			p := Point{r, c, k}
			heap.Push(xs, p)
		}
	}

	return -1
}

type Point struct {
	row, col int
	k        int
}

type PointHeap []Point

func (p PointHeap) Len() int           { return len(p) }
func (p PointHeap) Less(i, j int) bool { return p[i].k < p[j].k }
func (p PointHeap) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p *PointHeap) Push(x any) {
	*p = append(*p, x.(Point))
}

func (p *PointHeap) Pop() any {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[0 : n-1]
	return x
}
