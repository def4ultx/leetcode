package main

import "fmt"

func main() {
	pre := [][]int{
		// {1, 0},

		{1, 0},
		{0, 1},

		// {1, 0},
		// {2, 1},
		// {3, 0},
		// {4, 1},
		// {4, 2},
		// {4, 3},
	}
	canFinish(2, pre)
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int][]int)
	degree := make([]int, numCourses)

	for _, edge := range prerequisites {
		parent, child := edge[1], edge[0]

		_, ok := graph[parent]
		if ok {
			graph[parent] = append(graph[parent], child)
		} else {
			graph[parent] = []int{child}
		}
		degree[child]++
	}
	// fmt.Println(graph)

	sources := make([]int, 0)
	for i, v := range degree {
		if v == 0 {
			sources = append(sources, i)
		}
	}

	fmt.Println(degree)
	fmt.Println(graph)

	completed := 0
	visited := make(map[int]struct{})
	for len(sources) != 0 {
		src := sources[0]
		sources = sources[1:]

		fmt.Println("visited", src, visited, degree)

		if _, ok := visited[src]; ok {
			continue
		}
		visited[src] = struct{}{}
		completed++

		fmt.Println(graph[src])
		for _, v := range graph[src] {
			degree[v]--
			fmt.Println(v, degree)
			if degree[v] == 0 {
				sources = append(sources, v)
			}
		}
	}

	fmt.Println("BFS complete")
	fmt.Println(completed, numCourses)
	return completed == numCourses
}
