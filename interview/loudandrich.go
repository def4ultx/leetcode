package main

import "fmt"

func main() {
	richer := [][]int{
		{1, 0},
		{2, 1},
		{3, 1},
		{3, 7},
		{4, 3},
		{5, 3},
		{6, 3},
	}
	quite := []int{3, 2, 5, 4, 6, 1, 7, 0}
	loudAndRich(richer, quite)
}

func loudAndRich(richer [][]int, quiet []int) []int {
	mapping := make(map[int]int)
	graph := make(map[int][]int)
	for _, v := range richer {
		dst, src := v[0], v[1]

		_, ok := graph[src]
		if ok {
			graph[src] = append(graph[src], dst)
		} else {
			graph[src] = []int{dst}
		}

		mapping[src]++
		mapping[dst]++
	}

	n := len(mapping)
	res := make([]int, n)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = []int{-1, -1}
	}

	for i := 0; i < n; i++ {
		_, idx := find(graph, i, quiet, dp)
		res[i] = idx
	}
	fmt.Println(res)
	return res
}

func find(graph map[int][]int, current int, quiet []int, dp [][]int) (int, int) {
	if dp[current][0] != -1 {
		return dp[current][0], dp[current][1]
	}
	fmt.Println("visit", current)
	edges := graph[current]

	min, idx := quiet[current], current
	for _, e := range edges {
		m, i := find(graph, e, quiet, dp)
		if m < min {
			min, idx = m, i
		}
	}
	dp[current][0] = min
	dp[current][1] = idx
	return min, idx
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
