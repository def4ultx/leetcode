package main

import (
	"fmt"
	"math"
)

func main() {
	// modifiedGraphEdges(5, [][]int{{4, 1, -1}, {2, 0, -1}, {0, 3, -1}, {4, 3, -1}}, 0, 1, 5)
	// modifiedGraphEdges(3, [][]int{{0, 1, -1}, {0, 2, 5}}, 0, 2, 6)
	// modifiedGraphEdges(4, [][]int{{0, 1, -1}, {1, 2, -1}, {3, 1, -1}, {3, 0, 2}, {0, 2, 5}}, 2, 3, 8)
	modifiedGraphEdges(5, [][]int{{1, 4, 1}, {2, 4, -1}, {3, 0, 2}, {0, 4, -1}, {1, 3, 10}, {1, 0, 10}}, 0, 2, 15)
	// modifiedGraphEdges(5, [][]int{{1, 4, 1}, {2, 4, -1}, {3, 0, 2}, {0, 4, -1}, {1, 3, 10}, {1, 0, 10}}, 0, 2, 5)
	// modifiedGraphEdges(4, [][]int{{0, 1, -1}, {2, 0, 2}, {3, 2, 6}, {2, 1, 10}, {3, 0, -1}}, 1, 3, 12)
}

type Pair struct {
	weight int
	index  int
}

func shortestPath(n int, graph [][]Pair, src, dst int) []int {
	seen := make([]bool, n)
	prev := make([]int, n)
	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt32
		prev[i] = -1
	}
	dist[src] = 0

	for {
		u := findMinIndex(dist, seen)
		if u == -1 || u == dst {
			break
		}
		seen[u] = true

		for v, w := range graph[u] {
			if seen[v] {
				continue
			}
			if w.weight == 0 {
				continue
			}
			w.weight = Abs(w.weight)

			alt := dist[u] + w.weight
			if alt < dist[v] {
				dist[v] = alt
				prev[v] = u
			}
		}
	}

	return dist
}

func modifiedShortestPath(n int, graph [][]Pair, src, dst int, distances []int, diff int, edges [][]int) []int {
	seen := make([]bool, n)
	prev := make([]int, n)
	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt32
		prev[i] = -1
	}
	dist[src] = 0

	for {
		u := findMinIndex(dist, seen)
		if u == -1 || u == dst {
			break
		}
		seen[u] = true

		for v, w := range graph[u] {
			if seen[v] {
				continue
			}
			if w.weight == 0 {
				continue
			}
			w.weight = Abs(w.weight)

			fmt.Println("process", u, v)

			if graph[u][v].weight == -1 {
				nw := diff + distances[v] - dist[u]
				if nw > w.weight {
					fmt.Println("update", nw, diff, distances[v], dist[u])
					edges[w.index][2] = nw
					w.weight = nw
				}
			}

			alt := dist[u] + w.weight
			if alt < dist[v] {
				dist[v] = alt
				prev[v] = u
			}
		}
	}

	fmt.Println("path", prev)
	return dist
}

func findMinIndex(dist []int, seen []bool) int {
	idx, max := -1, math.MaxInt32
	for i, v := range dist {
		if !seen[i] && v < max {
			idx = i
			max = v
		}
	}
	return idx
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func modifiedGraphEdges(n int, edges [][]int, source int, destination int, target int) [][]int {

	graph := make([][]Pair, n)
	for i := range graph {
		graph[i] = make([]Pair, n)
	}

	for i, v := range edges {
		src, dst, weight := v[0], v[1], v[2]

		graph[src][dst] = Pair{weight, i}
		graph[dst][src] = Pair{weight, i}
	}

	dist := shortestPath(n, graph, source, destination)
	diff := target - dist[destination]
	if diff < 0 {
		return nil
	}

	fmt.Println(dist, diff)
	fmt.Println("done 1sp")
	dist = modifiedShortestPath(n, graph, source, destination, dist, diff, edges)

	if dist[destination] != target {
		return nil
	}

	for i := range edges {
		if edges[i][2] == -1 {
			edges[i][2] = 1
		}
	}

	fmt.Println("result")
	fmt.Println(edges)
	return edges
}

// type Edge struct {
// 	src, dst, weight int
// }

// type Heap []Edge

// // index := source
// // for {
// // 	i := prev[index]
// // }

// if true {

// 	for i := range graph {
// 		fmt.Println(graph[i])
// 	}
// }

// // fmt.Println(distance)
// // fmt.Println(dist)
// // fmt.Println(prev)

// path := []int{destination}
// index := destination
// for {
// 	i := prev[index]
// 	if i == -1 {
// 		break
// 	}

// 	path = append(path, i)
// 	index = i
// }

// edge1, edge2 := -1, -1

// modified := false
// diff := target - dist[destination]
// for i := len(path) - 1; i > 0; i-- {
// 	if graph[path[i]][path[i-1]] == -1 {
// 		modified = true
// 		graph[path[i]][path[i-1]] = diff + 1
// 		edge1, edge2 = path[i], path[i-1]
// 		// fmt.Println("update", graph[path[i]][path[i+1]])
// 		break
// 	}
// }

// fmt.Println(graph)
// fmt.Println(diff, modified)

// if diff != 0 && modified == false {
// 	return [][]int{}
// }

// if edge1 > edge2 {
// 	edge1, edge2 = edge2, edge1
// }

// fmt.Println(edge1, edge2)

// for i := range edges {
// 	a, b := edges[i][0], edges[i][1]
// 	if a > b {
// 		a, b = b, a
// 	}

// 	if a == edge1 && b == edge2 {
// 		edges[i][2] = diff + 1
// 	}
// 	if edges[i][2] == -1 {
// 		edges[i][2] = 1
// 	}
// }
