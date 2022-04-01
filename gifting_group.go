package main

import "fmt"

// expected 2
// 110
// 110
// 001

// expected 2
// 1100
// 1110
// 0110
// 0001

// expected 5
// 10000
// 01000
// 00100
// 00010
// 00001

func main() {
	xs := []string{"110", "110", "001"}
	// xs := []string{"1100", "1110", "0110", "0001"}
	// xs := []string{"10000", "01000", "00100", "00010", "00001"}

	countGroups(xs)
}

type Pair struct {
	sender   int
	receiver int
}

func countGroups(related []string) int32 {
	// Write your code here

	size := len(related)

	matrix := make([][]int, size)
	for i, str := range related {
		matrix[i] = make([]int, size)
		for j, ch := range str {
			if ch == '0' {
				matrix[i][j] = 0
			} else {
				matrix[i][j] = 1
			}
		}
	}
	// fmt.Println(matrix)

	known := make(map[int][]int)
	for i := 0; i < size; i++ {
		known[i] = make([]int, 0)
	}

	mapping := make(map[Pair]struct{})
	for i := range matrix {
		for j := range matrix[i] {
			if i == j {
				continue
			}
			if matrix[i][j] == 0 {
				continue
			}
			p := Pair{
				sender:   Min(i, j),
				receiver: Max(i, j),
			}
			fmt.Println(p)
			mapping[p] = struct{}{}

			known[i] = append(known[i], j)
			known[j] = append(known[j], i)
		}
	}
	fmt.Println(mapping)

	visited := make([]bool, size)
	count := 0
	for p := range mapping {
		if visited[p.sender] {
			continue
		}
		if visited[p.receiver] {
			continue
		}

		markVisited(visited, p.sender, known)
		markVisited(visited, p.receiver, known)
		count++
	}

	fmt.Println("count", count)
	fmt.Println(visited)

	for _, v := range visited {
		if !v {
			count++
		}
	}
	fmt.Println("final", count)
	return int32(count)
}

func markVisited(visited []bool, current int, known map[int][]int) {
	if visited[current] {
		return
	}

	visited[current] = true

	for _, v := range known[current] {
		markVisited(visited, v, known)
	}
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
