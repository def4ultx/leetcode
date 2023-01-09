package main

import (
	"fmt"
	"sort"
)

func main() {
	carFleet(12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3})
	carFleet(10, []int{6, 8}, []int{3, 2})
}

func carFleet(target int, position []int, speed []int) int {
	if len(position) <= 1 {
		return len(position)
	}

	type Entry struct {
		position  int
		speed     int
		remaining float32
	}

	entries := make([]Entry, len(position))
	for i := 0; i < len(position); i++ {
		entries[i] = Entry{position[i], speed[i], 0}

	}
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].position > entries[j].position
	})

	fmt.Println(entries)

	entries[0].remaining = float32(target-entries[0].position) / float32(entries[0].speed)
	stack := []Entry{entries[0]}

	for i := 1; i < len(entries); i++ {
		dist := target - entries[i].position
		remaining := float32(dist) / float32(entries[i].speed)

		top := stack[len(stack)-1]
		fmt.Println(remaining, top.remaining)
		if remaining > top.remaining {
			entries[i].remaining = remaining
			stack = append(stack, entries[i])
		}
	}

	fmt.Println(stack)

	return len(stack)
}
