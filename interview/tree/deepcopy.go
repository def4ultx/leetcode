package main

import "fmt"

func main() {
	n7 := &Node{7, nil, nil}
	n13 := &Node{13, nil, nil}
	n11 := &Node{11, nil, nil}
	n10 := &Node{10, nil, nil}
	n1 := &Node{1, nil, nil}

	n7.Next = n13
	n13.Next = n11
	n11.Next = n10
	n10.Next = n1

	n7.Random = nil
	n13.Random = n7
	n11.Random = n1
	n10.Random = n11
	n1.Random = n11

	copyRandomList(n7)
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	mapping := make(map[*Node]*Node)

	fmt.Println(head, head.Next, head.Next.Next, head.Next.Next.Next, head.Next.Next.Next.Next)

	dummy := &Node{}
	curr, temp := head, dummy
	for curr != nil {
		n := &Node{
			Val:    curr.Val,
			Next:   nil,
			Random: curr.Random,
		}

		temp.Next = n
		temp = temp.Next

		mapping[curr] = n
		curr = curr.Next
	}

	fmt.Println(dummy.Next, dummy.Next.Next, dummy.Next.Next.Next, dummy.Next.Next.Next.Next, dummy.Next.Next.Next.Next.Next)

	curr, temp = dummy.Next, head
	for curr != nil {
		curr.Random = mapping[temp.Random]
		curr, temp = curr.Next, temp.Next
	}

	fmt.Println(dummy.Next, dummy.Next.Next, dummy.Next.Next.Next, dummy.Next.Next.Next.Next, dummy.Next.Next.Next.Next.Next)
	fmt.Println(dummy.Next.Random, dummy.Next.Next.Random, dummy.Next.Next.Next.Random, dummy.Next.Next.Next.Next.Random, dummy.Next.Next.Next.Next.Next.Random)

	r := dummy.Next
	dummy.Next = nil
	return r
}
