package main

import (
	"container/heap"
	"fmt"
)

func main() {
	swimInWater([][]int{{0, 1}, {2, 3}})
}

func swimInWater(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	pq := &PathHeap{}
	heap.Init(pq)

	pq.Push(Path{0, 0, grid[0][0]})
	for heap.Interface.Len(pq) > 0 {
		p := heap.Pop(pq).(Path)

		fmt.Println("visited", p, p.row, m-1, p.col, n-1, p.row == m-1 && p.col == n-1)

		if p.row == m-1 && p.col == n-1 {
			fmt.Println("found result", p)
			return p.Cost
		}
		if p.row < 0 || p.col < 0 || p.row >= m || p.col >= n {
			continue
		}
		if visited[p.row][p.col] {
			continue
		}
		visited[p.row][p.col] = true

		directions := []struct{ i, j int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		for _, dir := range directions {
			r, c := p.row+dir.i, p.col+dir.j
			if r < 0 || c < 0 || r >= m || c >= n {
				continue
			}
			cost := Max(p.Cost, grid[r][c])
			pp := Path{r, c, cost}

			heap.Push(pq, pp)
		}
	}
	return 0
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Path struct {
	row, col int
	Cost     int
}

type PathHeap []Path

func (p PathHeap) Len() int           { return len(p) }
func (p PathHeap) Less(i, j int) bool { return p[i].Cost < p[j].Cost }
func (p PathHeap) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p *PathHeap) Push(x interface{}) {
	*p = append(*p, x.(Path))
}

func (p *PathHeap) Pop() interface{} {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[0 : n-1]
	return x
}
