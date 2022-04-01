package main

import (
	"sort"
	"strconv"
	"strings"
)

func main() {

	logs := []string{"88 99 200", "88 99 300", "99 32 100", "12 12 15"}
	threshold := 2
	processLogs(logs, int32(threshold))
}

func processLogs(logs []string, threshold int32) []string {
	// Write your code here

	mapping := make(map[int]int)
	for _, v := range logs {
		xs := strings.Split(v, " ")

		sender, _ := strconv.Atoi(xs[0])
		receiver, _ := strconv.Atoi(xs[1])

		if sender != receiver {
			mapping[sender]++
			mapping[receiver]++
		} else {
			mapping[sender]++
		}
	}

	th := int(threshold)
	arr := make([]int, 0)
	for k, v := range mapping {
		if v >= th {
			arr = append(arr, k)
		}
	}
	sort.Ints(arr)

	xs := make([]string, 0)
	for _, v := range arr {
		xs = append(xs, strconv.Itoa(v))
	}

	return xs
}
