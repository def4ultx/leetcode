package main

import (
	"fmt"
	"sort"
)

func main() {
	edges := [][]int{
		{0, 3}, {0, 4}, {1, 3}, {2, 4}, {2, 7}, {3, 5}, {3, 6}, {3, 7}, {4, 6},
	}
	getAncestors(8, edges)
}

func getAncestors(n int, edges [][]int) [][]int {
	children := make([][]int, n)
	for i := range children {
		children[i] = make([]int, 0)
	}
	degree := make([]int, n)
	for _, e := range edges {
		parent, child := e[0], e[1]
		children[parent] = append(children[parent], child)
		degree[child]++
	}
	fmt.Println(children)
	fmt.Println(degree)

	queue := make([]int, 0)
	for i, v := range degree {
		if v == 0 {
			queue = append(queue, i)
		}
	}

	ancestor := make([]map[int]bool, n)
	for i := range ancestor {
		ancestor[i] = make(map[int]bool)
	}

	for len(queue) > 0 {
		parent := queue[0]
		queue = queue[1:]

		for _, c := range children[parent] {
			ancestor[c][parent] = true
			for k := range ancestor[parent] {
				ancestor[c][k] = true
			}

			degree[c]--
			if degree[c] == 0 {
				queue = append(queue, c)
			}
		}
	}

	fmt.Println(ancestor)

	result := make([][]int, n)
	for i := range ancestor {
		for k := range ancestor[i] {
			result[i] = append(result[i], k)
		}
		sort.Ints(result[i])

	}

	fmt.Println(result)
	return result
}
