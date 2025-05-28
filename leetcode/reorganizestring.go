package main

import (
	"container/heap"
	"fmt"
)

func main() {
	reorganizeString("aab")
	reorganizeString("abcd")
	reorganizeString("aaab")
	reorganizeString("aaaaabbcc")
	reorganizeString("aaaaaabbcc")
}

func reorganizeString(s string) string {

	count := make(map[rune]int)
	for _, ch := range s {
		count[ch]++
	}

	var total, max int
	for _, v := range count {
		if v > max {
			max = v
		}
		total += v
	}

	if max > total-max+1 {
		return ""
	}

	pq := &ItemHeap{}
	heap.Init(pq)
	for k, v := range count {
		item := Item{k, v}
		heap.Push(pq, item)
	}

	result := make([]rune, 0)

	var prev Item
	for pq.Len() > 0 {
		pop := heap.Pop(pq).(Item)
		// fmt.Println("pop", pop)

		result = append(result, pop.Ch)
		pop.Count--

		if prev.Count > 0 {
			heap.Push(pq, prev)
		}
		prev = pop

		if pq.Len() == 0 && pop.Count > 0 {
			return ""
		}
	}

	r := string(result)
	fmt.Println("result", r)
	return r
}

type Item struct {
	Ch    rune
	Count int
}

type ItemHeap []Item

func (h ItemHeap) Len() int           { return len(h) }
func (h ItemHeap) Less(i, j int) bool { return h[i].Count > h[j].Count }
func (h ItemHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *ItemHeap) Push(x any) {
	*h = append(*h, x.(Item))
}

func (h *ItemHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
