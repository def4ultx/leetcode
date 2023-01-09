package main

import (
	"container/heap"
	"fmt"
)

func main() {
	// assignTasks([]int{3, 3, 2}, []int{1, 2, 3, 2, 1, 2})          // [2,2,0,2,1,2]
	// assignTasks([]int{5, 1, 4, 3, 2}, []int{2, 1, 2, 4, 5, 2, 1}) // [1,4,1,4,1,3,2]
	// assignTasks([]int{1, 1}, []int{100, 100, 100, 100, 2}) // [0 1 0 1 0]
	assignTasks([]int{1, 2, 3}, []int{5, 4, 3, 1, 2}) // 0, 1, 2, 0, 1
}

func assignTasks(servers []int, tasks []int) []int {

	available := &ServerHeap{}
	busy := &TaskHeap{}

	heap.Init(available)
	heap.Init(busy)

	for i, v := range servers {
		s := Server{
			weight: v,
			index:  i,
		}
		heap.Push(available, s)
	}

	result := make([]int, len(tasks))

	time, idx := 0, 0
	for idx < len(tasks) {
		// pop all available server from busy server
		// pick available server
		// assign task with available server to busy heap

		for {
			t := busy.Peek()
			if t == nil {
				break
			}
			task := t.(Task)
			if task.finished > time {
				// maybe need to set current time = task.finished?
				if available.Len() == 0 {
					time = task.finished
				}
				break
			}

			heap.Pop(busy)
			result[task.index] = task.server
			server := Server{
				weight: servers[task.server],
				index:  task.server,
			}
			heap.Push(available, server)
		}

		// if available.Len() <= 0 {
		// 	// hmm?
		// 	// time++
		// 	// fmt.Println("possible?!?", available, busy, time)
		// 	continue
		// }

		for available.Len() > 0 && time >= idx && idx < len(tasks) {
			server := heap.Pop(available).(Server)
			task := Task{
				finished: time + tasks[idx],
				index:    idx,
				server:   server.index,
			}
			heap.Push(busy, task)
			idx++
		}

		// fmt.Println(time, idx, available.Len(), "--", time+tasks[idx])
		// server := heap.Pop(available).(Server)
		// task := Task{
		// 	finished: time + tasks[idx],
		// 	index:    idx,
		// 	server:   server.index,
		// }
		// heap.Push(busy, task)

		// idx++
		// if !canAssignMultipleTask {
		// 	time++
		// }
		time++
	}

	// drain heap in busy server
	for busy.Len() > 0 {
		task := heap.Pop(busy).(Task)
		result[task.index] = task.server
	}

	fmt.Println(result)

	return result
}

type Server struct {
	weight int
	index  int
}

type ServerHeap []Server

func (h ServerHeap) Len() int      { return len(h) }
func (h ServerHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h ServerHeap) Less(i, j int) bool {
	if h[i].weight == h[j].weight {
		return h[i].index < h[j].index
	}
	return h[i].weight < h[j].weight
}

func (h *ServerHeap) Push(x interface{}) {
	*h = append(*h, x.(Server))
}

func (h *ServerHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Task struct {
	finished int
	index    int
	server   int
}

type TaskHeap []Task

func (h TaskHeap) Len() int           { return len(h) }
func (h TaskHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h TaskHeap) Less(i, j int) bool { return h[i].finished < h[j].finished }

func (h *TaskHeap) Push(x interface{}) {
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
	old := *h
	if len(old) == 0 {
		return nil
	}
	return old[0]
}
