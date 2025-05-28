package main

import (
	"fmt"
)

func main() {
	n1 := &TreeNode{1, nil, nil}
	n2 := &TreeNode{2, nil, nil}
	n3 := &TreeNode{3, nil, nil}
	n4 := &TreeNode{4, nil, nil}
	n5 := &TreeNode{5, nil, nil}
	n6 := &TreeNode{6, nil, nil}
	n7 := &TreeNode{7, nil, nil}
	// n8 := &TreeNode{8, nil, nil}

	n1.Left = n3
	n1.Right = n4
	n3.Right = n2

	n4.Left = n6
	n4.Right = n5
	n5.Right = n7
	treeQueriesLevel(n1, []int{3, 4})
	treeQueriesLevel(n1, []int{3})
	// [4] == 2

	// [5,4,6,2,null,null,8,1,3,7], [8] == 3

	// 		    5
	// 	   4           6
	// 	2,  null    null,8
	//  1,3     7

	// n5.Left = n4
	// n5.Right = n6
	// n4.Left = n2
	// n6.Right = n8
	// n2.Left = n1
	// n2.Right = n3
	// n8.Left = n7

	// treeQueriesLevel(n5, []int{8})
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func treeQueriesLevel(root *TreeNode, queries []int) []int {
	n := 0
	depth, height := [100001]int{}, [100001]int{}

	var dfs func(*TreeNode, int) int
	dfs = func(root *TreeNode, level int) int {
		if root == nil {
			return 0
		}
		n = Max(n, root.Val)

		fmt.Println("visit", root.Val, level)

		depth[root.Val] = level
		l := dfs(root.Left, level+1) + 1
		r := dfs(root.Right, level+1) + 1
		height[root.Val] = Max(l, r)

		return Max(l, r)
	}
	h := dfs(root, 0)
	n += 1

	fmt.Println(depth[:n])
	fmt.Println(height[:n])
	fmt.Println("len", n)

	type Tuple struct {
		Height int
		Val    int
	}
	mapping := make([][2]Tuple, h)
	sortTuple := func(a, b, c Tuple) (Tuple, Tuple) {
		if c.Height > a.Height {
			return c, a
		}
		if c.Height > b.Height {
			return a, c
		}
		return a, b
	}

	for i := 1; i < n; i++ {
		// mapping[depth[i]] = append(mapping[depth[i]], Tuple{height[i], i})
		first, second, tuple := mapping[depth[i]][0], mapping[depth[i]][1], Tuple{height[i], i}
		mapping[depth[i]][0], mapping[depth[i]][1] = sortTuple(first, second, tuple)
	}
	for i, xs := range mapping {
		// 	sort.Slice(xs, func(i, j int) bool { return xs[i].Height > xs[j].Height })
		fmt.Println(i, xs)
	}

	result := make([]int, len(queries))
	for i, v := range queries {
		fmt.Println("i,v", i, v, "at", depth[v])
		d := depth[v]
		// if len(mapping[d]) == 1 {
		if mapping[d][1].Val == 0 {
			fmt.Println("none")
			result[i] = d - 1
		} else if mapping[d][0].Val == v {
			fmt.Println("mapping")
			result[i] = mapping[d][1].Height + d - 1
		} else {
			fmt.Println("else")
			result[i] = mapping[d][0].Height + d - 1
		}
	}
	fmt.Println("result", result)
	return result
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func treeQueries(root *TreeNode, queries []int) []int {

	mapping := make(map[int]*PrefixSumNode)
	prefix, _ := copyTree(nil, root, mapping)

	result := make([]int, len(queries))
	for i, v := range queries {
		result[i] = calculateNewHeight(prefix, mapping[v]) - 1
	}
	fmt.Println("result", result)
	return result
}

func calculateNewHeight(root *PrefixSumNode, removed *PrefixSumNode) int {
	current := removed
	parent := removed.Parent

	height := 0
	for parent != nil {
		if parent.Left == current {
			height = Max(height, parent.RightHeight) + 1
		} else {
			height = Max(height, parent.LeftHeight) + 1
		}
		parent, current = parent.Parent, parent
	}

	return height
}

// func calculateNewHeight(root *PrefixSumNode, removed *PrefixSumNode) int {

// 	current := removed

// 	path := make([]*PrefixSumNode, 0)
// 	for ; current != nil; current = current.Parent {
// 		path = append(path, current)
// 		fmt.Println(current.Node.Val)
// 	}
// 	fmt.Println("path", path, removed)

// 	height := Max(path[len(path)-1].LeftHeight, path[len(path)-1].RightHeight)
// 	for i := len(path) - 1; i > 0; i-- {
// 		fmt.Println("walk", path[i].Node.Val, path[i].LeftHeight, path[i].RightHeight)
// 		if path[i].Left == path[i-1] && path[i].RightHeight >= path[i].LeftHeight {
// 			return height
// 		}
// 		if path[i].Right == path[i-1] && path[i].LeftHeight >= path[i].RightHeight {
// 			return height
// 		}
// 	}

// 	height = height - Abs(path[1].LeftHeight-path[1].RightHeight)
// 	// var other int
// 	// if path[len(path)-2] == root.Left {
// 	// 	other = root.RightHeight
// 	// } else {
// 	// 	other = root.LeftHeight
// 	// }

// 	return 0
// }

func copyTree(parent *PrefixSumNode, curr *TreeNode, mapping map[int]*PrefixSumNode) (*PrefixSumNode, int) {
	if curr == nil {
		return nil, 0
	}

	n := &PrefixSumNode{
		Node:   curr,
		Parent: parent,
	}
	mapping[curr.Val] = n
	n.Left, n.LeftHeight = copyTree(n, curr.Left, mapping)
	n.Right, n.RightHeight = copyTree(n, curr.Right, mapping)

	return n, Max(n.LeftHeight, n.RightHeight) + 1
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type PrefixSumNode struct {
	Node        *TreeNode
	Parent      *PrefixSumNode
	Left        *PrefixSumNode
	Right       *PrefixSumNode
	LeftHeight  int
	RightHeight int
}
