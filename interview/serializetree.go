package main

import (
	"strconv"
	"strings"
)

func main() {

}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "nil,"
	}
	return strconv.Itoa(root.Val) + "," + this.serialize(root.Left) + this.serialize(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	xs := strings.Split(data, ",")

	idx := 0
	var deserialize func() *TreeNode
	deserialize = func() *TreeNode {
		if xs[idx] == "nil" {
			idx++
			return
		}
		n, _ := strconv.Atoi(xs[idx])
		idx++
		root := &TreeNode{n, nil, nil}
		root.Left = deserialize()
		root.Right = deserialize()
		return root
	}
	return deserialize()
}
