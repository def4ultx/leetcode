package main

import (
	"fmt"
	"math"
)

// this problem is reduce to shortest path with exactly k edge
// algo
// dist = [] with inf
// for i = 0...k
// 		temp = [] with inf
// 		for e := edge
//			temp[e.dst] = min(temp[e.dst], tmp[e.src] + w)
//		dist = temp
// return dist[dst] if not inf

func main() {
	findCheapestPrice(4, [][]int{
		{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200},
	}, 0, 3, 1)

	findCheapestPrice(4, [][]int{
		{0, 1, 2}, {1, 2, 1}, {2, 0, 10}, {2, 3, 15},
	}, 1, 3, 0)
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {

	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[src] = 0

	for i := 0; i <= k; i++ {
		// for i := 0; i < k; i++ {
		temp := make([]int, n)
		for j := range temp {
			temp[j] = math.MaxInt32
		}
		temp[src] = 0

		for _, v := range flights {
			s, d, w := v[0], v[1], v[2]
			fmt.Println(i, s, d, w, "|", temp[d], dist[s]+w, Min(temp[d], dist[s]+w))
			temp[d] = Min(temp[d], dist[s]+w)
		}

		fmt.Println(dist, temp)

		for j := range dist {
			dist[j] = temp[j]
		}
	}

	fmt.Println("final")
	fmt.Println(dist)

	if dist[dst] == math.MaxInt32 {
		return -1
	}
	return dist[dst]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
