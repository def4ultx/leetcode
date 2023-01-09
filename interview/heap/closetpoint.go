package main

import (
	"container/heap"
	"fmt"
)

func main() {
	data := [][]int{
		{3, 3}, {5, -1}, {-2, 4},
	}
	kClosest(data, 2)
}

func kClosest(points [][]int, k int) [][]int {
	h := &Heap{}
	heap.Init(h)

	for i, v := range points {
		dist := (-v[0])*(-v[0]) + (-v[1])*(-v[1])
		point := Point{
			index:    i,
			distance: dist,
		}

		top, ok := h.Peek().(Point)
		fmt.Println("loop, point", point, ok, top)
		if h.Len() < k {
			fmt.Println("push point", point)
			heap.Push(h, point)
		} else if ok && point.distance < top.distance {
			fmt.Println("push point2", point)
			heap.Push(h, point)
			if h.Len() > k {
				heap.Pop(h)
			}
		}
	}
	fmt.Println(h, h.Len())

	result := make([][]int, k)
	size := h.Len()
	for i := 0; i < size; i++ {
		fmt.Println("i", i)
		point := heap.Pop(h).(Point)

		v := points[point.index]
		result[i] = []int{v[0], v[1]}
	}
	fmt.Println(result)

	return result
}

type Point struct {
	index    int
	distance int
}

type Heap []Point

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].distance > h[j].distance }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(Point))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *Heap) Peek() interface{} {
	old := *h
	if len(old) == 0 {
		return nil
	}
	return old[0]
}
