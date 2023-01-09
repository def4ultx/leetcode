package main

import (
	"fmt"
	"testing"
)

func TestFindImbalanceGroup(t *testing.T) {
	testcases := []struct {
		input    []int
		expected int
	}{

		{[]int{4, 1, 3, 2}, 3},
		{[]int{2, 0, 1}, 1},
		{[]int{1, 0, 2}, 1},
		{[]int{44, 11, 33, 22}, 10},
		{[]int{11, 22, 33, 44}, 10},
		{[]int{44, 33, 22, 11}, 10},
		// {[]int{44, 33, 22, 11, 55}, 20},
		// {[]int{44, 11, 33, 22, 55}, 20},
	}

	for i, tc := range testcases {
		tcName := fmt.Sprintf("tc #%d", i)
		t.Run(tcName, func(t *testing.T) {
			got := findImbalance(tc.input)
			if got != tc.expected {
				t.Errorf("got %v, want %v", got, tc.expected)
			}
		})
	}
}
