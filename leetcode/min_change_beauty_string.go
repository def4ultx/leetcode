package main

func minChanges(s string) int {
	var count int
	for i := 0; i < len(s); i += 2 {
		if s[i] != s[i+1] {
			count++
		}
	}
	return count
}
