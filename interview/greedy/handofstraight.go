package main

import (
	"container/heap"
	"fmt"
)

func main() {
	isNStraightHand([]int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3)
}

func isNStraightHand(hand []int, groupSize int) bool {

	h := &IntHeap{}
	heap.Init(h)

	if len(hand)%groupSize != 0 {
		return false
	}

	set := make(map[int]int)
	for _, v := range hand {
		set[v]++
	}
	for k := range set {
		heap.Push(h, k)
	}
	fmt.Println(hand)
	fmt.Println(set)

	count := 0
	index := 0
	for count < len(hand) {
		// v := hand[index]

		v := h.Top().(int)
		for i := v; i < v+groupSize; i++ {
			fmt.Println(i, v, set[i], set, count, index)
			_, ok := set[i]
			if !ok {
				fmt.Println("not ok")
				return false
			}
			// if c == 0 {
			// 	fmt.Println("exhaust")
			// 	return false
			// }
			set[i]--

			count++
			if set[i] == 0 {
				index++
				if i != h.Top().(int) {
					return false
				}
				heap.Pop(h)
			}
		}
	}
	return true
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap) Top() interface{} {
	x := *h
	return x[0]
}
