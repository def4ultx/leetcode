package main

func main() {

}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if isSameTree(root, subRoot) {
		return true
	}
	if root == nil {
		return false
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSameTree(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	if t1.Val != t2.Val {
		return false
	}
	return isSameTree(t1.Left, t2.Left) || isSameTree(t1.Right, t2.Right)
}

// func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
// 	t := merkle(subRoot)
// 	s := merkle(root)
// }

// func dfs(root *TreeNode) bool {
// 	if root == nil {
// 		return false
// 	}
// 	return
// }

// func hash(s string) uint32 {
// 	h := fnv.New32a()
// 	h.Write([]byte(s))
// 	return h.Sum32()
// }

// func merkle(root *TreeNode) uint32 {
// 	if root == nil {
// 		return 0
// 	}

// 	left := merkle(root.Left)
// 	right := merkle(root.Right)
// 	m := hash(strconv.Itoa(root.Val))

// 	return left + m + right
// }
