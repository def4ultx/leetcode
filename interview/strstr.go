package main

import (
	"fmt"
	"hash/fnv"
)

func main() {
	// KMP("hello", "ll")
	// KMP("hellosssssssssssssssssss", "aabaabaaa")
	// KMP("hellosssssssssssssssssss", "dsgwadsqz")
	// KMP("hellosssssssssssssssssss", "aaacaaaa")
	// KMP("hellosssssssssssssssssss", "acaacaaaa")
	// KMP("mississippi", "issip")
	KMP("mississippi", "issipi")
}

func strStr(haystack string, needle string) int {
	return -1
}

func KMP(haystack, needle string) int {

	table := make([]int, len(needle))
	i, j := 0, 1
	for j < len(needle) {
		if needle[j] == needle[i] {
			table[j] = i + 1
			i++
			j++
		} else {
			if i == 0 {
				table[j] = 0
				j++
			} else {
				fmt.Println("set table", j, i, table, table[i-1])
				i = table[i-1]
			}
		}
	}
	fmt.Println(table)

	i, idx := 0, 0
	for i < len(haystack) {
		if haystack[i] == needle[idx] {
			idx++
			i++
		} else {
			if idx == 0 {
				i++
			} else {
				idx = table[idx-1]
			}
		}

		fmt.Println(i, idx, len(needle))

		if idx == len(needle) {
			return i - idx + 1
		}
	}
	return -1
}

func Rolling(haystack, needle string) int {
	temp := hash(needle)
	for i := 0; i <= len(haystack)-len(needle); i++ {
		h := hash(haystack[i : i+len(needle)])
		if h == temp {
			return i
		}
	}
	return -1
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
