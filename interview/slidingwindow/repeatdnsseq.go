package main

import "fmt"

func main() {
	findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT")
}

func findRepeatedDnaSequences(s string) []string {
	if len(s) < 10 {
		return []string{}
	}

	dna := s[:10]
	mapping := make(map[string]int)
	mapping[dna] = 0

	temp := make(map[string]struct{})
	for i := 1; i < len(s)-9; i++ {
		dna = dna[1:]
		dna += string(s[i+9])

		// idx, found := mapping[dna]
		// fmt.Println("new dna", dna, i, idx, found)
		_, found := mapping[dna]
		if found {
			temp[dna] = struct{}{}
		} else {
			mapping[dna] = i
		}
	}

	result := make([]string, 0)
	for k := range temp {
		result = append(result, k)
	}
	fmt.Println(result)
	return result
}
