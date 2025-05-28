package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	// [[1,4],[2,3],[4,6]], targetFriend = 1 // 1
	// [[3,10],[1,5],[2,6]], targetFriend = 0 // 2
	// [33889,98676],[80071,89737],[44118,52565],[52992,84310],[78492,88209],[21695,67063],[84622,95452],[98048,98856],[98411,99433],[55333,56548],[65375,88566],[55011,62821],[48548,48656],[87396,94825],[55273,81868],[75629,91467], 6 // 2
	// fmt.Println(smallestChair([][]int{{1, 4}, {2, 3}, {4, 6}}, 1))
	// fmt.Println(smallestChair([][]int{{3, 10}, {1, 5}, {2, 6}}, 0))
	// fmt.Println(smallestChair([][]int{{33889, 98676}, {80071, 89737}, {44118, 52565}, {52992, 84310}, {78492, 88209}, {21695, 67063}, {84622, 95452}, {98048, 98856}, {98411, 99433}, {55333, 56548}, {65375, 88566}, {55011, 62821}, {48548, 48656}, {87396, 94825}, {55273, 81868}, {75629, 91467}}, 6))
	fmt.Println(smallestChair([][]int{{2, 30},
		{17, 24},
		{3, 6},
		{7, 19},
		{16, 22},
		{1, 14},
		{20, 27},
		{28, 31},
		{29, 32},
		{10, 11},
		{13, 23},
		{8, 12},
		{4, 5},
		{21, 26},
		{9, 18},
		{15, 25}}, 6))
}

func smallestChair(times [][]int, targetFriend int) int {
	type Tuple struct {
		Index int
		Time  int
	}

	n := len(times)
	arrivals := make([]Tuple, n)
	leavings := make([]Tuple, n)
	for i, v := range times {
		arrivals[i] = Tuple{i, v[0]}
		leavings[i] = Tuple{i, v[1]}
	}

	sort.Slice(arrivals, func(i, j int) bool { return arrivals[i].Time < arrivals[j].Time })
	sort.Slice(leavings, func(i, j int) bool { return leavings[i].Time < leavings[j].Time })

	// available := &Heap{}
	available := make(Heap, n)
	for i := range available {
		available[i] = i
	}
	heap.Init(&available)
	used := make([]int, n)

	// fmt.Println(arrivals, n)
	// fmt.Println(leavings, n)

	// count := 1
	// mm := make(map[int]int)
	// left, right := 0, 0
	// for left < n || right < n {
	// 	if left >= n || leavings[right].Time <= arrivals[left].Time {
	// 		fmt.Println(left, right, n, leavings[right].Time, count)
	// 		mm[leavings[right].Time] = count
	// 		right++
	// 	} else {
	// 		fmt.Println(left, right, n, arrivals[left].Time, count)
	// 		mm[arrivals[left].Time] = count
	// 		left++
	// 	}
	// 	count++
	// }
	// for _, v := range times {
	// 	fmt.Println(mm[v[0]], ",", mm[v[1]])
	// }

	fmt.Println(arrivals)
	fmt.Println(leavings)

	left, right := 0, 0
	for left < n || right < n {
		if left >= n-1 || leavings[right].Time <= arrivals[left].Time {
			chair := used[leavings[right].Index]

			fmt.Println("right", right, leavings[right], "return chair", chair)

			heap.Push(&available, chair)
			right++
		} else {
			if arrivals[left].Index == targetFriend {
				return heap.Pop(&available).(int)
			}

			chair := heap.Pop(&available).(int)

			fmt.Println("left", left, arrivals[left], "take chair", chair)

			used[arrivals[left].Index] = chair
			left++
		}
	}

	return -1
}

type Heap []int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] < h[j] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func smallestChair2(times [][]int, targetFriend int) int {
	n := len(times)

	h := &TupleHeap{} // TODO: Change this to array
	for i := range times {
		heap.Push(h, Tuple{times[i][0], i + 1})
		heap.Push(h, Tuple{times[i][1], -(i + 1)})
	}

	// available := &Heap{}
	available := make(Heap, n)
	for i := range available {
		available[i] = i
	}
	heap.Init(&available)
	used := make([]int, n+1)

	for h.Len() > 0 {
		pop := heap.Pop(h).(Tuple)
		if pop.Index == targetFriend+1 {
			return heap.Pop(&available).(int)
		}

		if pop.Index > 0 {
			chair := heap.Pop(&available).(int)
			used[pop.Index] = chair
		} else {
			heap.Push(&available, used[pop.Index])
		}
	}

	return -1
}

type Tuple struct {
	Time  int
	Index int
}

type TupleHeap []Tuple

func (h TupleHeap) Len() int           { return len(h) }
func (h TupleHeap) Less(i, j int) bool { return h[i].Time < h[j].Time }
func (h TupleHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *TupleHeap) Push(x any) {
	*h = append(*h, x.(Tuple))
}

func (h *TupleHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
