package main

import (
	"fmt"
	"sort"
)

func main() {
	tickets := [][]string{
		{"MUC", "LHR"},
		{"JFK", "MUC"},
		{"SFO", "SJC"},
		{"LHR", "SFO"},
	}
	// tickets := [][]string{
	// 	{"JFK", "SFO"},
	// 	{"JFK", "ATL"},
	// 	{"SFO", "ATL"},
	// 	{"ATL", "JFK"},
	// 	{"ATL", "SFO"},
	// }
	// tickets := [][]string{
	// 	{"JFK", "SFO"},
	// 	{"SFO", "JFK"},
	// 	{"JFK", "ATL"},
	// }
	findItinerary2(tickets)
}

// euler circuit vs Hamiltonian path => visit vertex at once
//

// Start with an empty stack and an empty circuit (eulerian path).
// - If all vertices have same out-degrees as in-degrees - choose any of them.
// - If all but 2 vertices have same out-degree as in-degree, and one of those 2 vertices has out-degree with one greater than its in-degree, and the other has in-degree with one greater than its out-degree - then choose the vertex that has its out-degree with one greater than its in-degree.
// - Otherwise no euler circuit or path exists.
// If current vertex has no out-going edges (i.e. neighbors) - add it to circuit, remove the last vertex from the stack and set it as the current one. Otherwise (in case it has out-going edges, i.e. neighbors) - add the vertex to the stack, take any of its neighbors, remove the edge between that vertex and selected neighbor, and set that neighbor as the current vertex.
// Repeat step 2 until the current vertex has no more out-going edges (neighbors) and the stack is empty.

// 1. Start with an empty stack and an empty circuit (eulerian path).
// - If all vertices have same out-degrees as in-degrees - choose any of them.
// - If all but 2 vertices have same out-degree as in-degree,
//		and one of those 2 vertices has out-degree with one greater than its in-degree,
//		and the other has in-degree with one greater than its out-degree
//		- then choose the vertex that has its out-degree with one greater than its in-degree.
// - Otherwise no euler circuit or path exists.
// 2. If current vertex has no out-going edges (i.e. neighbors)
// - add it to circuit, remove the last vertex from the stack and set it as the current one.
// Otherwise (in case it has out-going edges, i.e. neighbors)
// - add the vertex to the stack,
//		take any of its neighbors,
//		remove the edge between that vertex and selected neighbor,
//		and set that neighbor as the current vertex.
// 3. Repeat step 2 until the current vertex has no more out-going edges (neighbors) and the stack is empty.

// Hierholzer's 1873 paper provides a different method for finding Euler cycles that is more efficient than Fleury's algorithm:
// 1. Choose any starting vertex v, and follow a trail of edges from that vertex until returning to v.
// It is not possible to get stuck at any vertex other than v, because the even degree of all vertices ensures that
// when the trail enters another vertex w there must be an unused edge leaving w. The tour formed in this way is a closed tour
// but may not cover all the vertices and edges of the initial graph.

// 2. As long as there exists a vertex u that belongs to the current tour but that has adjacent edges not part of the tour,
// start another trail from u, following unused edges until returning to u, and join the tour formed in this way to the previous tour.

// 3. Since we assume the original graph is connected, repeating the previous step will exhaust all edges of the graph.

// Choose vertex v
// Create stack, result, append v into stack, set current = v
// while len(stack) > 0
// get edge from current node
// if len(edge) > 0
// pop edge from the list, append current to stack, set current = edge that just pop from list
// else, append current into result, pop vertex from stack, set current = vertex

func findItinerary(tickets [][]string) []string {

	// create graph
	// sort edge in the graph
	// create euler circuit
	// return the circuit

	in, out := make(map[string]int), make(map[string]int)
	for _, route := range tickets {
		out[route[0]]++
		in[route[1]]++
	}

	graph := make(map[string][]string)
	for _, route := range tickets {
		fmt.Println(route)
		src, dst := route[0], route[1]

		_, ok := graph[src]
		if ok {
			graph[src] = append(graph[src], dst)
		} else {
			graph[src] = []string{dst}
		}
	}
	for _, v := range graph {
		sort.Strings(v)
	}
	// fmt.Println(graph)

	// if all v have same in/out degree, pick any of them
	// if all but 2 have same in/out, those two must 1 node with have 1 out > in, and 1 in > out
	// pick 1 out > in
	// else no euler path

	result := make([]string, 0)

	stack := make([]string, 0)
	stack = append(stack, "JFK")

	current := "JFK"
	for len(stack) > 0 {
		edges := graph[current]
		if len(edges) == 0 {
			result = append(result, current)

			current = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, current)
			next := edges[0]
			graph[current] = graph[current][1:]

			current = next
		}
	}

	Reverse(result)
	fmt.Println("result")
	fmt.Println(result)
	return result
}

func Reverse(xs []string) {
	n := len(xs)
	for i := 0; i < len(xs)/2; i++ {
		xs[i], xs[n-i-1] = xs[n-i-1], xs[i]
	}
}

func findItinerary2(tickets [][]string) []string {

	graph := make(map[string][]string)
	for _, route := range tickets {
		fmt.Println(route)
		src, dst := route[0], route[1]

		_, ok := graph[src]
		if ok {
			graph[src] = append(graph[src], dst)
		} else {
			graph[src] = []string{dst}
		}
	}
	for _, v := range graph {
		sort.Strings(v)
	}
	fmt.Println("graph", graph)

	result := make([]string, 0)
	result = append(result, "JFK")
	dfs(graph, "JFK", len(tickets), &result)

	fmt.Println("result")
	fmt.Println(result)
	return result
}

func dfs(graph map[string][]string, node string, ticket int, result *[]string) bool {
	edge, ok := graph[node]
	fmt.Println("visit", node, edge, ok, result)
	if !ok {
		fmt.Println("not found", result)
		return false
	}
	if len(*result) == ticket+1 {
		fmt.Println("return", result)
		return true
	}

	// fmt.Println("result", result)

	list := make([]string, len(edge))
	copy(list, edge)
	for i, v := range list {
		// remove v
		// add add v to result
		// dfs(v)
		// if dfs == true, return
		// else pop v from result
		// add dst to edges

		graph[node] = remove(graph[node], i)
		*result = append(*result, v)

		found := dfs(graph, v, ticket, result)
		if found {
			return true
		}
		old := *result
		*result = old[:len(old)-1]

		graph[node] = insert(graph[node], i, v)
	}
	return false
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func insert(slice []string, i int, v string) []string {
	r := append(slice[:i], v)
	r = append(r, slice[i:]...)
	return r
}
