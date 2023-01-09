package main

import (
	"fmt"
)

func main() {

	n := 7
	edges := [][]int{
		{0, 1},
		{0, 2},
		{1, 4},
		{1, 5},
		{2, 3},
		{2, 6},
	}
	// hasApples := []bool{false, false, true, false, true, true, false}
	// hasApples := []bool{false, false, true, false, false, true, false}
	// hasApples := []bool{false, false, false, false, false, false, false}
	hasApples := []bool{true, false, false, false, false, false, false}

	// n := 8
	// edges := [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 4}, {2, 5}, {2, 6}, {4, 7}}
	// hasApples := []bool{true, true, false, true, false, true, true, false}

	// n := 4
	// edges := [][]int{{0, 2}, {0, 3}, {1, 2}}
	// hasApples := []bool{false, true, false, false}

	minTime(n, edges, hasApples)
}

type Node struct {
	nodes    []*Node
	value    int
	hasApple bool
}

func minTime(n int, edges [][]int, hasApple []bool) int {
	nodes := make([]*Node, 0)
	for i := 0; i < n; i++ {
		node := &Node{
			nodes:    make([]*Node, 0),
			value:    i,
			hasApple: hasApple[i],
		}
		nodes = append(nodes, node)
	}

	for _, v := range edges {
		src, dst := nodes[v[0]], nodes[v[1]]
		src.nodes = append(src.nodes, dst)
		dst.nodes = append(dst.nodes, src)
	}

	// fmt.Println(nodes)
	// for _, v := range nodes {
	// 	left, right := -1, -1
	// 	if v.left != nil {
	// 		left = v.left.value
	// 	}
	// 	if v.right != nil {
	// 		right = v.right.value
	// 	}
	// 	fmt.Println(v.value, left, right)
	// }
	// fmt.Println("----")

	root := nodes[0]
	visited := make(map[int]struct{})
	cost := findMin(root, visited)
	if cost >= 2 {
		cost -= 2
	}
	fmt.Println(cost)

	return cost
}

func findMin(root *Node, visited map[int]struct{}) int {
	if root == nil {
		return 0
	}
	if _, ok := visited[root.value]; ok {
		return 0
	}

	visited[root.value] = struct{}{}

	// var left, right int
	// if root.left != nil {
	// 	left = findMin(root.left)
	// }
	// if root.right != nil {
	// 	right = findMin(root.right)
	// }
	var total int
	for _, node := range root.nodes {
		total += findMin(node, visited)
	}

	if total > 0 || root.hasApple {
		total += 2
	}

	// fmt.Println(root.value, left, right, root.hasApple, cost)
	return total
}
