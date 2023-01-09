package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	minCostConnectPointsOptimize([][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}})
	// minCostConnectPointsOptimize([][]int{{3, 12}, {-2, 5}, {-4, 1}})
}

type Node struct {
	src int
	key int
}

// Algorithm
// 1) Create a set mstSet that keeps track of vertices already included in MST.
// 2) Assign a key value to all vertices in the input graph. Initialize all key values as INFINITE. Assign key value as 0 for the first vertex so that it is picked first.
// 3) While mstSet doesn’t include all vertices
// ….a) Pick a vertex u which is not there in mstSet and has minimum key value.
// ….b) Include u to mstSet.
// ….c) Update key value of all adjacent vertices of u. To update the key values, iterate through all adjacent vertices. For every adjacent vertex v, if weight of edge u-v is less than the previous key value of v, update the key value as weight of u-v
// The idea of using key values is to pick the minimum weight edge from cut. The key values are used only for vertices which are not yet included in MST, the key value for these vertices indicate the minimum weight edges connecting them to the set of vertices included in MST.

// Algorithm - Optimize

// Initialize some variables:

// nn - Number of nodes of the graph.
// mstCostmstCost - Cost to build the MST.
// edgesUsededgesUsed - Number of edges included in the MST.
// inMSTinMST - Array to track which nodes are already part of the MST.
// minDistminDist - Array to track the minimum edge weight to reach the i^{th}i
// th
//   node from any node that is already in the tree.
// Initially, we start with node 00, and the cost to reach this node will be 00. To signify this, we set minDist[0]minDist[0] equal to 00.

// We will try adding nodes to our MST until edgesUsededgesUsed becomes equal to nn.

// We pick the node which uses the lowest weight edge and is not present in the MST.
// We increment edgesUsededgesUsed by 11, mark this node as included in the MST, and add the edge weight used to reach this node to the mstCostmstCost.
// Try updating the minimum distance to all adjacent nodes in minDistminDist.
// We return the total cost of the MST, mstCostmstCost.

func minCostConnectPointsOptimize(points [][]int) int {
	size := len(points)
	if size < 2 {
		return 0
	}

	visited := make([]bool, size)
	distance := make([]int, size)
	for i := range distance {
		distance[i] = math.MaxInt32
	}
	distance[0] = 0

	index, connected, total := 0, 0, 0
	for connected < size {
		if visited[index] {
			continue
		}
		visited[index] = true

		connected++
		total += distance[index]

		min, minIndex := math.MaxInt32, -1
		for i := 0; i < size; i++ {
			dist := Abs(points[index][0]-points[i][0]) + Abs(points[index][1]-points[i][1])

			fmt.Println("set min index", "|", !visited[i], distance[i], dist, "index", i)

			if !visited[i] {
				if distance[i] > dist {
					distance[i] = dist
				}
				if distance[i] < min {
					fmt.Println("set to", i, "distance", distance[i])
					min, minIndex = distance[i], i
				}
			}
		}
		if minIndex != -1 {
			index = minIndex
		}
	}

	fmt.Println(total)
	return total
}

func minCostConnectPoints(points [][]int) int {
	nodes := make([]Node, len(points))
	for i := range points {
		nodes[i] = Node{
			src: i,
			key: 0,
		}
	}

	minHeap := &Heap{}
	heap.Init(minHeap)

	// for i := 1; i < len(points); i++ {
	// 	heap.Push(minHeap, nodes[i])
	// }
	heap.Push(minHeap, nodes[0])

	fmt.Println(minHeap)

	visited := make([]bool, len(points))
	connected, minCost := 0, 0
	for minHeap.Len() > 0 && connected < len(points) {
		node := heap.Pop(minHeap).(Node)
		fmt.Println(node, visited, minHeap)

		if visited[node.src] {
			continue
		}
		visited[node.src] = true

		minCost += node.key
		connected++

		src := node.src
		for i := 0; i < len(points); i++ {
			// if i == src {
			// 	continue
			// }
			if visited[i] {
				continue
			}
			cost := Abs(points[src][0]-points[i][0]) + Abs(points[src][1]-points[i][1])
			fmt.Println(src, i, cost)
			heap.Push(minHeap, Node{i, i, cost})
		}
	}

	fmt.Println(minCost)
	return minCost
}

func Abs(x int) int {
	if x < 0 {
		return 0 - x
	}
	return x
}

type Heap []Node

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].key < h[j].key }
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
