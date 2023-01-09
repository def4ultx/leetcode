package main

func main() {

}

func longestSubstring(s string, k int) int {
	count := make([]int, 26)
	for _, v := range s {
		count[v-'a']++
	}

	for i, c := range s {
		if count[c-'a'] >= k {
			continue
		}
		a := longestSubstring(s[0:i], k)
		b := longestSubstring(s[i+1:], k)
		return Max(a, b)
	}
	return len(s)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
