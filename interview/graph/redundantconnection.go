package main

import "fmt"

func main() {
	findRedundantConnection([][]int{
		{1, 2},
		{1, 3},
		{2, 3},
	})
}

func findRedundantConnection(edges [][]int) []int {
	// Can you reach dst from src in each item in slice
	// if yes, dont add, else add it to the graph
	dsu := &DisjointUnionSet{
		parent: make([]int, 1001),
	}
	for i := range dsu.parent {
		dsu.parent[i] = -1
	}

	for _, e := range edges {
		ok := dsu.Union(e[0], e[1])
		if !ok {
			return e
		}
	}
	return []int{}
}

type DisjointUnionSet struct {
	parent []int
}

func (dsu *DisjointUnionSet) Find(x int) int {
	if dsu.parent[x] == -1 {
		return x
	}
	return dsu.Find(dsu.parent[x])
}

func (dsu *DisjointUnionSet) Union(x, y int) bool {
	px, py := dsu.Find(x), dsu.Find(y)
	fmt.Println(px, py)
	if px == py {
		fmt.Println("return false", x, y)
		return false
	}
	dsu.parent[px] = py
	return true
}
