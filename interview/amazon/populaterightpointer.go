package main

import "fmt"

func main() {
	n1 := &Node{1, nil, nil, nil}
	n2 := &Node{2, nil, nil, nil}
	n3 := &Node{3, nil, nil, nil}
	n4 := &Node{4, nil, nil, nil}
	n5 := &Node{5, nil, nil, nil}
	n6 := &Node{6, nil, nil, nil}
	n7 := &Node{7, nil, nil, nil}

	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
	n3.Left = n6
	n3.Right = n7

	connect(n1)
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

type LevelNode struct {
	*Node
	level int
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := make([]LevelNode, 0)
	queue = append(queue, LevelNode{root, 0})

	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]

		if top.Node == nil {
			continue
		}
		fmt.Println("visit", top.Node, top.level)

		var next LevelNode
		if len(queue) > 0 {
			next = queue[0]
		}

		if next.Node != nil && top.level == next.level {
			top.Next = next.Node
		}
		l := LevelNode{top.Left, top.level + 1}
		r := LevelNode{top.Right, top.level + 1}

		queue = append(queue, l, r)

	}
	return root
}
