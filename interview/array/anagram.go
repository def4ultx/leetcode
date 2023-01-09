package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println('a', 'A')
	groupAnagramsV3([]string{"bbb", "eat", "tea", "tan", "ate", "nat", "bat", "aaa"})
}

// n^2
func groupAnagramsV3(strs []string) [][]string {

	const CHAR_OFFSET = 97

	mapping := make(map[string][]string)
	for i, word := range strs {

		freq := make([]int, 26)
		for _, ch := range word {
			freq[ch-CHAR_OFFSET]++
		}

		freqStr := make([]string, 26)
		for i, v := range freq {
			n := strconv.Itoa(v)
			freqStr[i] = n
		}

		key := strings.Join(freqStr, "#")
		if _, ok := mapping[key]; ok {
			mapping[key] = append(mapping[key], strs[i])
		} else {
			mapping[key] = []string{strs[i]}
		}
	}

	idx := 0
	result := make([][]string, len(mapping))
	for _, v := range mapping {
		result[idx] = v
		idx++
	}
	return result
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

// n*klogk
func groupAnagramsV2(strs []string) [][]string {

	words := make([]string, len(strs))
	for i, v := range strs {
		words[i] = SortString(v)
	}

	mapping := make(map[string][]string)
	for i, v := range words {
		if _, ok := mapping[v]; ok {
			mapping[v] = append(mapping[v], strs[i])
		} else {
			mapping[v] = []string{strs[i]}
		}
	}

	idx := 0
	result := make([][]string, len(mapping))
	for _, v := range mapping {
		result[idx] = v
		idx++
	}
	return result
}

// n^2*k
func groupAnagrams(strs []string) [][]string {
	isAnagrams := make([]int, len(strs))
	for i := range isAnagrams {
		isAnagrams[i] = -1
	}

	var index int
	for i, v := range strs {
		for j := i; j < len(strs); j++ {
			if isAnagrams[i] != -1 && isAnagrams[j] != -1 {
				continue
			}

			match := isAnagram(v, strs[j])
			if match {
				group := isAnagrams[i]
				if group == -1 {
					isAnagrams[i] = index
					isAnagrams[j] = index
					index++
				} else {
					isAnagrams[j] = isAnagrams[i]
				}
			}
		}
	}

	fmt.Println(isAnagrams)
	result := make([][]string, index)
	for i, v := range isAnagrams {
		if len(result[v]) == 0 {
			result[v] = make([]string, 0)
		}
		result[v] = append(result[v], strs[i])
	}

	fmt.Println(result)

	return result
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mapping := make(map[rune]int)

	for _, r := range s {
		mapping[r]++
	}

	for _, r := range t {
		count, ok := mapping[r]
		if !ok || count == 0 {
			return false
		}
		mapping[r]--
	}
	return true
}
