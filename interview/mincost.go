package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {

	minCostConnectPointsoptimize([][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}})
}

func minCostConnectPointsoptimize(points [][]int) int {
	dist := make([]int, len(points))
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[0] = 0
	visited := make([]bool, len(points))

	index, total, connected := 0, 0, 0
	for connected < len(points) {
		fmt.Println("visit", index)
		if visited[index] {
			continue
		}
		visited[index] = true

		connected++
		total += dist[index]

		idx, min := -1, math.MaxInt32
		for i := range points {
			if visited[i] {
				continue
			}

			cost := Abs(points[index][0]-points[i][0]) + Abs(points[index][1]-points[i][1])
			fmt.Println(index, i, cost)
			if cost < dist[i] {
				dist[i] = cost
			}
			if dist[i] < min {
				fmt.Println("set to", i, "distance", dist[i])
				min, idx = dist[i], i
			}
		}
		if idx != -1 {
			index = idx
		}
		break
	}
	return total
}

func minCostConnectPoints(points [][]int) int {
	ph := &PointHeap{}
	p0 := Point{0, 0}
	heap.Init(ph)
	heap.Push(ph, p0)

	total, connected := 0, 0
	visited := make(map[int]bool)
	for ph.Len() > 0 && connected != len(points)+1 {
		top := heap.Pop(ph).(Point)
		if visited[top.src] {
			continue
		}
		visited[top.src] = true
		connected++
		total += top.cost

		for i, p := range points {
			if visited[i] {
				continue
			}
			cost := Abs(p[0]-points[top.src][0]) + Abs(p[1]-points[top.src][1])
			pp := Point{i, cost}
			heap.Push(ph, pp)
		}
	}
	return total
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

type Point struct {
	src  int
	cost int
}

type PointHeap []Point

func (h PointHeap) Len() int           { return len(h) }
func (h PointHeap) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h PointHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PointHeap) Push(x interface{}) {
	*h = append(*h, x.(Point))
}

func (h *PointHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
