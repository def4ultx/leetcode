package main

import (
	"container/heap"
	"fmt"
)

func main() {
	f := MedianFinderConstructor()
	// f.AddNum(1)
	// f.AddNum(2)
	// f.FindMedian()
	// f.AddNum(3)
	// f.FindMedian()
	// f.AddNum(4)
	// f.AddNum(5)
	// f.AddNum(6)
	// f.AddNum(7)
	// f.FindMedian()

	f.AddNum(-1)
	f.FindMedian()
	f.AddNum(-2)
	f.FindMedian()
	f.AddNum(-3)
	f.FindMedian()
	f.AddNum(-4)
	f.FindMedian()
	f.AddNum(-5)
	f.FindMedian()
}

func logHeap(m MedianFinder) {
	fmt.Println("min", m.minHeap)
	fmt.Println("max", m.maxHeap)
}

type MedianFinder struct {
	minHeap Heap
	maxHeap Heap
}

func MedianFinderConstructor() MedianFinder {

	h, pq := &IntHeap{}, &IntPriorityQueue{}
	heap.Init(h)
	heap.Init(pq)

	return MedianFinder{
		minHeap: h,
		maxHeap: pq,
	}
}

func (this *MedianFinder) AddNum(num int) {
	top := this.maxHeap.Top()
	if top == nil {
		heap.Push(this.maxHeap, num)
		return
	}

	topMax := top.(int)
	fmt.Println("top?", topMax, "num?", num)
	fmt.Println(this.maxHeap)
	fmt.Println(this.minHeap)
	if num > topMax {
		fmt.Println("add to min heap")
		heap.Push(this.minHeap, num)
	} else {
		fmt.Println("add to max heap")
		heap.Push(this.maxHeap, num)
	}
	fmt.Println(this.maxHeap)
	fmt.Println(this.minHeap)

	if this.maxHeap.Len()-this.minHeap.Len() > 1 {
		x := heap.Pop(this.maxHeap).(int)
		fmt.Println("balance from left to right", x)
		heap.Push(this.minHeap, x)
	} else if this.minHeap.Len()-this.maxHeap.Len() > 1 {
		x := heap.Pop(this.minHeap).(int)
		fmt.Println("balance from right to left", x)
		heap.Push(this.maxHeap, x)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	var result float64

	fmt.Println("------")
	fmt.Println(this.maxHeap)
	fmt.Println(this.minHeap)

	if this.maxHeap.Len() == this.minHeap.Len() {
		l, r := this.maxHeap.Top().(int), this.minHeap.Top().(int)
		result = (float64(l) + float64(r)) / 2
	} else if this.maxHeap.Len() > this.minHeap.Len() {
		result = float64(this.maxHeap.Top().(int))
	} else {
		result = float64(this.minHeap.Top().(int))
	}

	fmt.Println("result", result)
	return result
}

type Heap interface {
	heap.Interface
	Top() interface{}
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

type IntPriorityQueue []int

func (h IntPriorityQueue) Len() int           { return len(h) }
func (h IntPriorityQueue) Less(i, j int) bool { return h[i] > h[j] }
func (h IntPriorityQueue) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntPriorityQueue) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntPriorityQueue) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntPriorityQueue) Top() interface{} {
	old := *h
	if len(old) == 0 {
		return nil
	}
	x := old[0]
	return x
}
