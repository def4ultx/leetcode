package main

import (
	"fmt"
)

func main() {
	// ms := map[int][]int{
	// 	1: {1, 2, 3},
	// }
	// fmt.Println(ms)
	// ms[1] = ms[1][1:]
	// fmt.Println(ms)
	// ms[1] = ms[1][1:]
	// fmt.Println(ms)
	// ms[1] = ms[1][1:]
	// fmt.Println(ms)
	// xs, ok := ms[1]
	// fmt.Println(xs, ok)
	// ms[1] = ms[1][1:]
	// fmt.Println(ms)
	// xs, ok = ms[1]
	// fmt.Println(xs, ok)

	pairs := [][]int{
		{5, 1},
		{4, 5},
		{11, 9},
		{9, 4},
	}
	validArrangement(pairs)
}

func validArrangement(pairs [][]int) [][]int {
	mapping := make(map[int][]int)
	nodes, in, out := make(map[int]struct{}), make(map[int]int), make(map[int]int)

	for _, p := range pairs {
		src, dst := p[0], p[1]
		edges, ok := mapping[src]
		if ok {
			mapping[src] = append(edges, dst)
		} else {
			mapping[src] = []int{dst}
		}
		in[dst]++
		out[src]++
		nodes[src] = struct{}{}
		nodes[dst] = struct{}{}
	}

	// find current
	// if all in,out = same, pick any of them
	// if all in,out = same except 2, pick node with out > in

	totalDegree := len(nodes)
	sameDegree := 0
	diffDegree := make([]int, 0)
	for k := range nodes {
		i, inOk := in[k]
		o, outOk := out[k]
		if inOk && outOk && i == o {
			sameDegree++
		} else {
			diffDegree = append(diffDegree, k)
		}
	}
	fmt.Println(in, out)
	fmt.Println(totalDegree, sameDegree, diffDegree)

	current := 0
	if totalDegree == sameDegree {
		for k := range nodes {
			current = k
			break
		}
	} else if len(diffDegree) == 2 {
		v1, v2 := diffDegree[0], diffDegree[1]

		if out[v1]-in[v1] == 1 && in[v2]-out[v2] == 1 {
			current = v1
		} else if out[v2]-in[v2] == 1 && in[v1]-out[v1] == 1 {
			current = v2
		} else {
			fmt.Println("here?")
			return [][]int{}
		}
	} else {
		fmt.Println("otherwise")
		return [][]int{}
	}

	path := make([]int, 0)
	stack := make([]int, 0)
	stack = append(stack, current)
	for len(stack) > 0 {
		// fmt.Println("visit")
		edges := mapping[current]
		if len(edges) > 0 {
			next := mapping[current][0]
			mapping[current] = mapping[current][1:]

			stack = append(stack, current)
			current = next
		} else {
			path = append(path, current)
			current = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
	}

	for i := 0; i < len(path)/2; i++ {
		path[i], path[len(path)-i-1] = path[len(path)-i-1], path[i]
	}
	result := make([][]int, 0)
	for i := 0; i < len(path)-1; i++ {
		result = append(result, []int{path[i], path[i+1]})
	}

	fmt.Println(result)
	return result
}
