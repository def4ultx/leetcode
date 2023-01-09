package main

import (
	"container/heap"
	"fmt"
)

func main() {

	// another solution
	// count letter appearance and store in hash[i]
	// find the letter with largest occurence.
	// put the letter into even index numbe (0, 2, 4 ...) char array
	// put the rest into the array

	// aab aaab aabc
	reorganizeString("aab")
	reorganizeString("aaab")
	reorganizeString("aabc")
	reorganizeString("aaabc")
}

func reorganizeString(s string) string {
	chars := make([]int, 26)
	for i := range s {
		idx := s[i] - 'a'
		chars[idx]++
	}

	pq := &CharPQ{}
	heap.Init(pq)
	for i, v := range chars {
		if v != 0 {
			c := Char{byte(i) + 'a', v, 0}
			heap.Push(pq, c)
		}
	}

	str := "-"
	for pq.Len() > 0 {
		top := heap.Pop(pq).(Char)

		lastChar := str[len(str)-1]
		if lastChar == top.char && pq.Len() == 0 {
			return ""
		}
		if lastChar == top.char {
			// pop another
			next := heap.Pop(pq).(Char)
			heap.Push(pq, top)
			top = next
		} else {
			// using current
		}

		str += string(top.char)
		top.count--
		top.lastUsed = len(str)
		if top.count > 0 {
			heap.Push(pq, top)
		}
	}

	fmt.Println(str)
	return str[1:]
}

type Char struct {
	char     byte
	count    int
	lastUsed int
}

type CharPQ []Char

func (h CharPQ) Len() int      { return len(h) }
func (h CharPQ) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h CharPQ) Less(i, j int) bool {
	if h[i].count == h[j].count {
		return h[i].lastUsed < h[j].lastUsed
	}
	return h[i].count > h[j].count
}
func (h *CharPQ) Push(x interface{}) {
	*h = append(*h, x.(Char))
}
func (h *CharPQ) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
