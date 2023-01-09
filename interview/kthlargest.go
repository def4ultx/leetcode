package main

import (
	"container/heap"
)

func main() {
	findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4)
}

func findKthLargest(nums []int, k int) int {
	pq := &IntHeap{}
	heap.Init(pq)

	for i := 0; i < k; i++ {
		heap.Push(pq, nums[i])
	}

	for i := k; i < len(nums); i++ {
		top := pq.Top().(int)
		if nums[i] >= top {
			heap.Push(pq, nums[i])
		}
		if pq.Len() > k {
			heap.Pop(pq)
		}
	}
	return heap.Pop(pq).(int)
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
