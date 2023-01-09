package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	tasks := [][]int{
		// {1, 2}, {2, 4}, {3, 2}, {4, 1},
		// {7, 10}, {7, 12}, {7, 5}, {7, 4}, {7, 2},
		{5, 2}, {7, 2}, {9, 4}, {6, 3}, {5, 10}, {1, 1},
	}
	getOrder(tasks)
}
func getOrder(tasks [][]int) []int {
	allTask := make([]Task, 0)
	for i, t := range tasks {
		t := Task{
			index:          i,
			enqueueAt:      t[0],
			processingTime: t[1],
			finishedAt:     -1,
		}
		allTask = append(allTask, t)
	}

	sort.Slice(allTask, func(i, j int) bool {
		if allTask[i].enqueueAt == allTask[j].enqueueAt {
			return allTask[i].processingTime < allTask[j].processingTime
		}
		return allTask[i].enqueueAt < allTask[j].enqueueAt
	})

	fmt.Println("all tasks", allTask)

	order := 0

	idx := 1
	currentTime := allTask[0].enqueueAt

	taskHeap := &TaskHeap{}
	heap.Init(taskHeap)
	taskHeap.Push(allTask[0])

	result := make([]int, 0)
	for taskHeap.Len() > 0 {
		fmt.Println("visit at time", currentTime, "with task", taskHeap)
		task := heap.Pop(taskHeap).(Task)
		if task.finishedAt == -1 {
			currentTime += task.processingTime
		} else {
			currentTime = task.finishedAt
		}

		fmt.Println("visit task", task, "at time", currentTime, "with task", taskHeap)
		fmt.Println("---")

		// result[task.index] = order
		result = append(result, task.index)
		order++

		for idx < len(tasks) {
			if allTask[idx].enqueueAt > currentTime {
				break
			}
			heap.Push(taskHeap, allTask[idx])
			idx++
		}

		if taskHeap.Len() == 0 && idx < len(tasks) {
			currentTime = allTask[idx].enqueueAt
			heap.Push(taskHeap, allTask[idx])
			idx++
		}
		fmt.Println(taskHeap, allTask, idx, currentTime)
	}

	fmt.Println(result)

	return result
}

type Task struct {
	index          int
	enqueueAt      int
	processingTime int
	finishedAt     int
}

type TaskHeap []Task

func (h TaskHeap) Len() int      { return len(h) }
func (h TaskHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h TaskHeap) Less(i, j int) bool {
	if h[i].processingTime == h[j].processingTime {
		return h[i].index < h[j].index
	}
	return h[i].processingTime < h[j].processingTime
}

func (h *TaskHeap) Push(x interface{}) {
	*h = append(*h, x.(Task))
}

func (h *TaskHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
