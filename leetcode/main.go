package main

import (
	"container/heap"
	"fmt"
)

func main() {
	// t := time.NewTimer(time.Second)
	// t.Reset(d time.Duration)

	n1 := &ListNode{1, nil}
	n2 := &ListNode{2, nil}
	n3 := &ListNode{3, nil}
	n4 := &ListNode{4, nil}

	n5 := &ListNode{5, nil}
	n6 := &ListNode{6, nil}
	n7 := &ListNode{7, nil}
	n8 := &ListNode{8, nil}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4

	n5.Next = n6
	n6.Next = n7
	n7.Next = n8

	lists := []*ListNode{n1, n5, nil}
	mergeKLists(lists)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	pq := make(PriorityQueue, 0)
	for _, v := range lists {
		if v != nil {
			pq = append(pq, v)
		}
	}
	heap.Init(&pq)

	head := &ListNode{}
	curr := head
	for pq.Len() > 0 {

		pop := heap.Pop(&pq).(*ListNode)

		fmt.Println("size", pq.Len(), pop)

		// if pop == nil {
		// 	continue
		// }

		if pop.Next != nil {
			heap.Push(&pq, pop.Next)
		}

		curr.Next = pop
		curr = curr.Next
	}

	return head.Next
}

type PriorityQueue []*ListNode

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(*ListNode))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}
