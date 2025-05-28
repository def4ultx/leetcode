package main

import "fmt"

func main() {
	// buildSequence("1234")
	// buildSequence("xyz")
	buildSequence("ba")
}

func buildSequence(str string) []string {
	results := make([]string, 0)
	temp := make([]byte, 0)

	buildSequenceRecursive(str, 0, &results, temp)

	fmt.Println("---result---")
	for _, v := range results {
		fmt.Println(v)
	}
	return results
}

func buildSequenceRecursive(str string, idx int, results *[]string, temp []byte) {
	if idx >= len(str) {
		if len(temp) != 0 {
			*results = append(*results, string(temp))
		}
		return
	}

	// fmt.Println("visit", idx)

	buildSequenceRecursive(str, idx+1, results, temp)

	if (len(temp) == 0) || (len(temp) > 0 && str[idx] > temp[len(temp)-1]) {
		temp = append(temp, str[idx])
		buildSequenceRecursive(str, idx+1, results, temp)
		temp = temp[:len(temp)-1]
	}
}
