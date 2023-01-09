package main

import (
	"fmt"
	"sort"
)

func main() {
	intervals := []*Interval{
		{0, 30},
		{5, 10},
		{15, 20},
	}
	fmt.Println(canAttendMeetings(intervals))

	intervals = []*Interval{
		{5, 8},
		{9, 15},
	}
	fmt.Println(canAttendMeetings(intervals))
}

type Interval struct {
	Start, End int
}

func canAttendMeetings(intervals []*Interval) bool {
	// Write your code here
	sort.Slice(intervals, func(i, j int) bool { return *&intervals[i].Start < *&intervals[j].Start })

	for i := 1; i < len(intervals); i++ {
		prev := *intervals[i-1]
		curr := *intervals[i]

		if prev.End > curr.Start {
			return false
		}
	}
	return true
}
