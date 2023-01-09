package main

import (
	"container/heap"
	"fmt"
)

func main() {
	leastInterval([]byte{'A', 'B', 'A'}, 2)                // 8
	leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2) // 8
	leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 0) // 6
	leastInterval([]byte{'A', 'A', 'A', 'A', 'A', 'A'}, 1) // 11
	leastInterval([]byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}, 2)
}

func leastInterval(tasks []byte, n int) int {

	// count all duplicate task, put to hashmap
	// add all task in hashmap with amount to heap
	// while heap is not empty
	// loop with n size, pop item and reduce amount by 1
	// add item size to total counter
	// push item back to heap
	// if there is a case where you cannot take item from heap due to same task to pull
	// get diff range and continue with diff range + nums task that can pull added to counter
	// if only 1/2 (< n) task left, add only lenght of 1/2 (< n) task

	counts := make(map[byte]int)
	for _, v := range tasks {
		counts[v]++
	}

	taskHeap := &TaskHeap{}
	heap.Init(taskHeap)

	for k, v := range counts {
		t := Task{k, v}
		heap.Push(taskHeap, t)
	}

	total := 0
	for taskHeap.Len() > 0 {
		local := 0
		seen := make([]Task, 0)
		for i := 0; i <= n; i++ {
			fmt.Println(i)
			local++

			if taskHeap.Len() == 0 {
				fmt.Println("empty task, idle")
				continue
			}

			t := heap.Pop(taskHeap)
			if t == nil {
				fmt.Println("should not exist")
				continue // no more item, add idle
			}

			task := t.(Task)
			fmt.Println("visit task", string(task.name), task.amount, local, total)
			if task.amount == 1 && taskHeap.Len() == 0 && len(seen) == 0 {
				fmt.Println("last task?")
				break // last task
			}
			if task.amount == 1 {
				continue // last take
			}

			task.amount--
			seen = append(seen, task)
		}

		for _, t := range seen {
			heap.Push(taskHeap, t)
		}

		total += local

		// push item back to heap
	}

	fmt.Println("total", total)

	return total
}

type Task struct {
	name   byte
	amount int
}

type TaskHeap []Task

func (h TaskHeap) Len() int           { return len(h) }
func (h TaskHeap) Less(i, j int) bool { return h[i].amount > h[j].amount }
func (h TaskHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *TaskHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Task))
}

func (h *TaskHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *TaskHeap) Peek() interface{} {
	x := *h
	return x[0]
}
