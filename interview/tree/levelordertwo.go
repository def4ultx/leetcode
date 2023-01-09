package main

func main() {

}

func levelOrderBottom(root *TreeNode) [][]int {
	type Node struct {
		*TreeNode
		level int
	}

	result := make([][]int, 0)

	queue := []Node{Node{root, 0}}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.TreeNode == nil {
			continue
		}

		if node.level >= len(result) {
			result = append(result, []int{})
		}
		result[node.level] = append(result[node.level], node.Val)

		queue = append(queue, Node{node.Left, node.level + 1})
		queue = append(queue, Node{node.Right, node.level + 1})
	}

	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-i-1] = result[len(result)-i-1], result[i]
	}

	return result
}
