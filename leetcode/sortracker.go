package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println("alps" < "bradford")

	tracker := Constructor()   // Initialize the tracker system.
	tracker.Add("bradford", 2) // Add location with name="bradford" and score=2 to the system.
	tracker.Add("branford", 3) // Add location with name="branford" and score=3 to the system.
	tracker.Get()              // The sorted locations, from best to worst, are: branford, bradford.
	// Note that branford precedes bradford due to its higher score (3 > 2).
	// This is the 1st time get() is called, so return the best location: "branford".
	tracker.Add("alps", 2) // Add location with name="alps" and score=2 to the system.
	tracker.Get()          // Sorted locations: branford, alps, bradford.
	// Note that alps precedes bradford even though they have the same score (2).
	// This is because "alps" is lexicographically smaller than "bradford".
	// Return the 2nd best location "alps", as it is the 2nd time get() is called.
	tracker.Add("orland", 2) // Add location with name="orland" and score=2 to the system.
	tracker.Get()            // Sorted locations: branford, alps, bradford, orland.
	// Return "bradford", as it is the 3rd time get() is called.
	tracker.Add("orlando", 3) // Add location with name="orlando" and score=3 to the system.
	tracker.Get()             // Sorted locations: branford, orlando, alps, bradford, orland.
	// Return "bradford".
	tracker.Add("alpine", 2) // Add location with name="alpine" and score=2 to the system.
	tracker.Get()            // Sorted locations: branford, orlando, alpine, alps, bradford, orland.
	// Return "bradford".
	tracker.Get() // Sorted locations: branford, orlando, alpine, alps, bradford, orland.
	// Return "orland".

	//["branford","alps","bradford","bradford","bradford","orland"]

	fmt.Println(tracker.h)

	tracker.Pop()
}

type SORTracker struct {
	h     *Heap
	count int
}

func Constructor() SORTracker {
	h := &Heap{}
	heap.Init(h)
	return SORTracker{h, 0}
}

func (this *SORTracker) Add(name string, score int) {
	heap.Push(this.h, Item{name, score})
}

func (this *SORTracker) Get() string {

	fmt.Println(len(*this.h), this.count, "-", *this.h)

	x := (*this.h)[this.count].Name
	this.count++

	fmt.Println(x)
	return x
}

func (this *SORTracker) Pop() {
	for this.h.Len() > 0 {
		fmt.Println(heap.Pop(this.h))
	}
}

type Item struct {
	Name  string
	Score int
}

type Heap []Item

func (h Heap) Len() int { return len(h) }
func (h Heap) Less(i, j int) bool {
	if h[i].Score == h[j].Score {
		return h[i].Name < h[j].Name
	}
	return h[i].Score > h[j].Score
}
func (h Heap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x any) {
	*h = append(*h, x.(Item))
}
func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
