package main

import (
	"fmt"
	"sort"
)

func main() {

	// pipes := []int{0, 3, 5}
	// end := 5
	// ei := sort.Search(len(pipes), func(i int) bool { return pipes[i] >= end })
	// if pipes[ei] != end {
	// 	ei--
	// }

	// start := 4
	// si := sort.Search(len(pipes), func(i int) bool { return pipes[i] >= start })
	// if pipes[si] != start {
	// 	si--
	// }

	// fmt.Println(ei, si)

	s := "|**|*|*"
	start := []int32{1, 1}
	end := []int32{5, 6}

	numberOfItems(s, start, end)
}

func numberOfItems(s string, startIndices []int32, endIndices []int32) []int32 {
	pipes := make([]int, 0)
	stars := 0

	mapping := make(map[int]int)

	for i, ch := range s {
		if ch == '|' {
			mapping[i] += stars
			pipes = append(pipes, i)
		} else if len(pipes) != 0 {
			stars++
		}
	}

	fmt.Println(pipes)
	fmt.Println(mapping)

	result := make([]int32, len(startIndices))
	for i := 0; i < len(startIndices); i++ {
		start, end := int(startIndices[i])-1, int(endIndices[i])-1

		// binary search
		si := sort.Search(len(pipes), func(i int) bool { return pipes[i] >= start })
		ei := sort.Search(len(pipes), func(i int) bool { return pipes[i] >= end })
		if pipes[si] != start {
			si--
		}
		if pipes[ei] != end {
			ei--
		}

		fmt.Println("si, ei")
		fmt.Println(si, ei)
		fmt.Println(si, pipes[si])
		fmt.Println(ei, pipes[ei])
		fmt.Println(mapping[pipes[ei]] - mapping[pipes[si]])
		result[i] = int32(mapping[pipes[ei]] - mapping[pipes[si]])

	}
	fmt.Println(result)

	return result

}

// type Pair struct {
// 	LastCompartmentIndex     int
// 	ItemSinceLastCompartment int
// 	TotalItem                int

// 	Start, End int
// }

// func numberOfItems(s string, startIndices []int32, endIndices []int32) []int32 {

// 	pairs := make([]Pair, 0)
// 	for i := 0; i < len(startIndices); i++ {
// 		p := Pair{
// 			LastCompartmentIndex: -1,
// 			Start:                int(startIndices[i]) - 1,
// 			End:                  int(endIndices[i]) - 1,
// 		}
// 		pairs = append(pairs, p)
// 	}

// 	mapping := make(map[int][]int)

// 	for idx, v := range pairs {
// 		for i := v.Start; i <= v.End; i++ {
// 			mapping[i] = append(mapping[i], idx)
// 		}
// 	}

// 	for i, ch := range s {
// 		pairIndexes := mapping[i]
// 		for index := range mapping {
// 			if ch == '|' {
// 				fmt.Println("compartment")
// 				if pairs[index].LastCompartmentIndex == -1 {
// 					pairs[index].LastCompartmentIndex = i
// 				} else {
// 					pairs[index].LastCompartmentIndex = i
// 					pairs[index].TotalItem += pairs[index].ItemSinceLastCompartment
// 					pairs[index].ItemSinceLastCompartment = 0
// 				}
// 			} else if ch == '*' {
// 				if pairs[index].LastCompartmentIndex != -1 {
// 					pairs[index].ItemSinceLastCompartment++
// 				}
// 			}
// 		}
// 		// for index := range pairs {
// 		// 	fmt.Println("start i", i, ch)
// 		// 	if pairs[index].Start <= i && pairs[index].End >= i {
// 		// if ch == '|' {
// 		// 	fmt.Println("compartment")
// 		// 	if pairs[index].LastCompartmentIndex == -1 {
// 		// 		pairs[index].LastCompartmentIndex = i
// 		// 	} else {
// 		// 		pairs[index].LastCompartmentIndex = i
// 		// 		pairs[index].TotalItem += pairs[index].ItemSinceLastCompartment
// 		// 		pairs[index].ItemSinceLastCompartment = 0
// 		// 	}
// 		// } else if ch == '*' {
// 		// 	if pairs[index].LastCompartmentIndex != -1 {
// 		// 		pairs[index].ItemSinceLastCompartment++
// 		// 	}
// 		// }
// 		// 	}
// 		// 	fmt.Println(pairs[index])
// 		// }
// 	}

// 	fmt.Println(pairs)

// 	result := make([]int32, 0)
// 	for _, v := range pairs {
// 		result = append(result, int32(v.TotalItem))
// 	}

// 	fmt.Println(result)
// 	// Write your code here
// 	return result
// }
