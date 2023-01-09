package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	// Input: times = {{2,1,1},{2,3,1},{3,4,1}}, n = 4, k = 2
	// Output: 2
	// Example 2:

	// networkDelayTime([][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2)
	// networkDelayTime([][]int{{1, 2, 1}}, 2, 1)
	// networkDelayTime([][]int{{1, 2, 1}}, 2, 2)
	// networkDelayTime([][]int{{1, 2, 1}, {2, 3, 2}, {1, 3, 4}}, 3, 1) // 3
	networkDelayTime([][]int{{2, 1, 5}, {2, 3, 5}, {3, 4, 5}, {1, 4, 3}}, 4, 2)
}

// Single source shortest path
// shortest path without stop condition
// shortests path
// dist = [] with inf
// dist[src] = 0
// init pq (using distance as identifier)
// add src to pq
// while pq > 0
//		pop pq
//		for e := edge
//			if dist[e.dst] > dist[e.src] + w
//				dist[e.dst] = dist[e.src] + w
//				add dst to pq

func networkDelayTime(times [][]int, n int, k int) int {

	graph := make(map[int][]Node)
	for _, v := range times {
		src, dst, weight := v[0]-1, v[1]-1, v[2]
		node := Node{dst, weight}

		if _, ok := graph[src]; ok {
			graph[src] = append(graph[src], node)
		} else {
			graph[src] = []Node{node}
		}
	}

	distances := make([]int, n)
	for i := range distances {
		distances[i] = math.MaxInt32
	}
	distances[k-1] = 0

	pq := &Heap{}
	heap.Init(pq)
	heap.Push(pq, Node{k - 1, 0})

	for pq.Len() > 0 {
		node := pq.Pop().(Node)
		fmt.Println("check this pq", pq)

		if node.weight > distances[node.index] {
			fmt.Println("check this")
			continue
		}

		nodes, ok := graph[node.index]
		if !ok {
			continue
		}
		fmt.Println("visited", node, nodes)

		for _, edge := range nodes {
			dist := edge.weight + node.weight
			fmt.Println("look edge", edge, dist, distances)
			if distances[edge.index] > dist {
				distances[edge.index] = dist
				heap.Push(pq, Node{edge.index, dist})
			}
		}
	}

	fmt.Println(distances)
	max := 0
	for _, v := range distances {
		if v > max {
			max = v
		}
		if v == math.MaxInt32 {
			return -1
		}
	}
	return max
}

type Node struct {
	index  int
	weight int
}

type Heap []Node

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].weight < h[j].weight }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(Node))
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

// // BFS with level and visit
// func networkDelayTime(times [][]int, n int, k int) int {
// 	type Node struct {
// 		index  int
// 		weight int
// 	}

// 	graph := make(map[int][]Node)
// 	for _, v := range times {
// 		src, dst, weight := v[0]-1, v[1]-1, v[2]
// 		node := Node{dst, weight}

// 		if _, ok := graph[src]; ok {
// 			graph[src] = append(graph[src], node)
// 		} else {
// 			graph[src] = []Node{node}
// 		}
// 	}

// 	fmt.Println("graph", graph)

// 	visited := make([]bool, n)
// 	queue := []Node{Node{k - 1, 0}}

// 	max := 0
// 	for len(queue) > 0 {
// 		node := queue[0]
// 		queue = queue[1:]

// 		fmt.Println("visited", node)
// 		if visited[node.index] {
// 			continue
// 		}
// 		visited[node.index] = true

// 		max = Max(max, node.weight)

// 		edges, ok := graph[node.index]
// 		if !ok {
// 			fmt.Println("not found graph")
// 			continue
// 		}

// 		for _, v := range edges {
// 			fmt.Println(v.index, v.weight)
// 			// if !visited[v.index] {
// 			// }
// 			queue = append(queue, Node{v.index, node.weight + v.weight})
// 		}
// 	}

// 	for _, v := range visited {
// 		if !v {
// 			fmt.Println("not possible")
// 			return -1
// 		}
// 	}

// 	fmt.Println(max)
// 	return max
// }

// func Max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
