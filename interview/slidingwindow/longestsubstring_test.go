package main

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	testcases := []struct {
		str  string
		want int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"au", 2},
		{"aab", 2},
		{"a", 1},
		{"", 0},
		{"dvdf", 3},
		{"tmmzuxt", 5},
		{"aabaab!bb", 3},
	}

	for _, tc := range testcases {
		t.Run(tc.str, func(t *testing.T) {
			got := lengthOfLongestSubstring(tc.str)
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
