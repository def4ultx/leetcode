package main

import "fmt"

func main() {
	pre := [][]int{
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

// time: o(v+e)
// space: o(v)
// Topological sort
// L <- empty list contain sorted nodes
// while exist node without permanent mark
// - select unmarked node
// - visit(n)
// def visit
//		if permanent mark(n) return
//		if tempo mark(n) return (not DAG)
//		mark n with tempo
//		visit all edge
//		mark n with permanent
//		add n to L

func findOrder(numCourses int, prerequisites [][]int) []int {

	graph := make(map[int][]int)
	degree := make([]int, numCourses)

	for _, edge := range prerequisites {
		parent, child := edge[0], edge[1]

		_, ok := graph[parent]
		if ok {
			graph[parent] = append(graph[parent], child)
		} else {
			graph[parent] = []int{child}
		}
		degree[parent]++
	}

	head := topologicalSort(numCourses, graph)
	fmt.Println("head", head)

	return head
}

func topologicalSort(numCourses int, graph map[int][]int) []int {
	head := []int{}

	marked := make([]int, numCourses)
	queue := make([]int, numCourses)
	for i := range queue {
		queue[i] = i
	}
	fmt.Println("queue", queue)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Println("while at node", node, marked[node], marked[node] != 0)
		if marked[node] == 2 {
			continue
		}

		isNotCycle := visit(node, graph, marked, &head)
		if !isNotCycle {
			fmt.Println("found cycle")
			return nil
		}
	}
	return head
}

func visit(node int, graph map[int][]int, marked []int, head *[]int) bool {
	fmt.Println("visit node", node, "with data", marked, head)

	// permanent mark
	if marked[node] == 2 {
		fmt.Println("permanent mark?")
		return true
	}

	// temporary mark, DAG
	if marked[node] == 1 {
		fmt.Println("temporary mark?")
		return false
	}

	marked[node] = 1
	edges := graph[node]
	for _, v := range edges {
		isNotCycle := visit(v, graph, marked, head)
		if !isNotCycle {
			fmt.Println("found cycle")
			return false
		}
	}

	marked[node] = 2
	*head = append(*head, node)
	return true
}
