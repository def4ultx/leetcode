package main

import (
	"container/heap"
)

func main() {
	// kSmallestPairs([]int{1, 7, 11}, []int{2, 4, 6}, 3)
	// kSmallestPairs([]int{1, 1, 2}, []int{1, 2, 3}, 2)
	// kSmallestPairs([]int{1, 2}, []int{3}, 3)
	kSmallestPairs([]int{1, 1, 2}, []int{1, 2, 3}, 10)
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {

	if len(nums1) == 0 || len(nums2) == 0 {
		return [][]int{}
	}

	h := &PairHeap{}
	p0 := Pair{nums1[0] + nums2[0], 0, 0}
	heap.Init(h)
	heap.Push(h, p0)

	// visit := make([][]bool, len(nums1))
	// for i := range visit {
	// 	visit[i] = make([]bool, len(nums2))
	// }
	visit := make(map[Pair]bool)

	result := make([][]int, 0)
	for k >= 0 && h.Len() > 0 {
		p := heap.Pop(h).(Pair)
		// fmt.Println("visit", p)
		if visit[p] {
			// fmt.Println("found visited", p)
			continue
		}
		visit[p] = true

		result = append(result, []int{nums1[p.l], nums2[p.r]})
		k--

		if p.l+1 < len(nums1) {
			pp := Pair{
				sum: nums1[p.l+1] + nums2[p.r],
				l:   p.l + 1,
				r:   p.r,
			}
			heap.Push(h, pp)
		}
		if p.r+1 < len(nums2) {
			pp := Pair{
				sum: nums1[p.l] + nums2[p.r+1],
				l:   p.l,
				r:   p.r + 1,
			}
			heap.Push(h, pp)
		}
	}
	return result

	// result := make([][]int, 0)
	// l, r := 0, 0
	// for i := 0; i < k && l < len(nums1) && r < len(nums2); i++ {
	// 	result = append(result, []int{nums1[l], nums2[r]})

	// 	fmt.Println(l, r)
	// 	if l == len(nums1)-1 {
	// 		r++
	// 	} else if r == len(nums2)-1 {
	// 		l++
	// 	} else {
	// 		left := nums1[l+1] + nums2[r]
	// 		right := nums1[l] + nums2[r+1]
	// 		if left < right {
	// 			l++
	// 		} else {
	// 			r++
	// 		}
	// 	}
	// }
	// fmt.Println(result)
	// return result

	// left, right := &IntHeap{}, &IntHeap{}
	// heap.Init(left)
	// heap.Init(right)

	// for _, v := range nums1 {
	// 	heap.Push(left, v)
	// }
	// for _, v := range nums2 {
	// 	heap.Push(right, v)
	// }

	// result := make([][]int, 0)
	// l := heap.Pop(left).(int)
	// r := heap.Pop(right).(int)
	// for i := 0; i < k; i++ {

	// }
	// return result
}

type Pair struct {
	sum  int
	l, r int
}

type PairHeap []Pair

func (h PairHeap) Len() int           { return len(h) }
func (h PairHeap) Less(i, j int) bool { return h[i].sum < h[j].sum }
func (h PairHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PairHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *PairHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
