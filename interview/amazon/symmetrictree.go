package main

func main() {

}

func isSymmetricIterative(root *TreeNode) bool {
	if root == nil {
		return true
	}

	type Pair struct {
		left, right *TreeNode
	}
	p0 := Pair{root.Left, root.Right}
	queue := make([]Pair, 0)
	queue = append(queue, p0)

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if p.left == nil && p.right == nil {
			continue
		} else if p.left == nil || p.right == nil {
			return false
		} else if p.left.Val != p.right.Val {
			return false
		}

		p1 := Pair{p.left.Left, p.right.Right}
		p2 := Pair{p.left.Right, p.right.Left}
		queue = append(queue, p1, p2)
	}
	return true
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return checkSymmetric(root.Left, root.Right)
}

func checkSymmetric(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}

	l := checkSymmetric(left.Left, right.Right)
	r := checkSymmetric(left.Right, right.Left)

	return l && r && (left.Val == right.Val)
}
