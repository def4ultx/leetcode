package main

import (
	"fmt"
	"sort"
)

// optimizing box weights

func main() {
	arr := []int32{3, 7, 5, 6, 2}
	minimalHeaviestSetA(arr)
}

func minimalHeaviestSetA(arr []int32) []int32 {
	// Write your code here

	xs := make([]int, 0)
	for _, v := range arr {
		xs = append(xs, int(v))
	}
	sort.Ints(xs)

	var total int
	for _, v := range xs {
		total += v
	}
	half := total / 2

	var count int
	var minIdx int
	for i := len(xs) - 1; i >= 0; i-- {
		if count > half {
			fmt.Println(i)
			minIdx = i + 1
			break
		}
		count += xs[i]
	}

	result := make([]int32, 0)
	for i := minIdx; i < len(xs); i++ {
		result = append(result, int32(xs[i]))
	}

	// fmt.Println(result)
	return result

	// h := &IntHeap{}
	// heap.Init(h)

	// arrs := make([]int, 0)
	// for _, v := range arr {
	// 	arrs = append(arrs, int(v))
	// }

	// var total int
	// for _, v := range arrs {
	// 	total += v
	// }
	// half := total / 2

	// for _, v := range arrs {
	// 	heap.Push(h, v)
	// }

	// xs := make([]int, 0)
	// var count int
	// for {
	// 	if count > half {
	// 		break
	// 	}
	// 	x := heap.Pop(h).(int)
	// 	xs = append(xs, x)

	// 	count += x
	// }

	// fmt.Println(xs)

	// sort.Ints(xs)

	// result := make([]int32, 0)
	// for _, v := range xs {
	// 	result = append(result, int32(v))
	// }
	// fmt.Println(result)

	// return result
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
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
