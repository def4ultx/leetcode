// Give a binary array, count minimum number of moves to get ones at one end and zeros at another end.
// example: [1,1,0,1,0,0] -> [1,1,1,0,0,0]
// [0,1,0,0,1,1] -> [0,0,0,1,1,1]

package main

import "fmt"

func main() {
	// countMin([]int{1, 1, 0, 1, 0, 0})                            // 1
	// countMin([]int{0, 1, 0, 0, 1, 1})                            // 1
	// countMin([]int{0, 1, 0, 1, 1, 0, 0, 1, 1, 0})                // 2
	// countMin([]int{0, 1, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0})    // 3
	countMin([]int{0, 1, 0, 1, 1, 0, 0, 0, 1, 1, 0, 0, 1, 1, 0}) // 3
}

func countMin(arr []int) int {
	size := len(arr)
	half := size / 2

	var limit int
	if size%2 == 0 {
		limit = half
	} else {
		limit = half + 1
	}

	var zero, one int
	for i := 0; i < limit; i++ {
		if arr[i] == 0 {
			zero++
		} else {
			one++
		}
	}

	var total int
	if zero < one && size%2 == 0 {
		total = half - one
	} else if zero >= one && size%2 == 0 {
		total = half - zero
	} else if zero < one && size%2 == 1 {
		total = limit - one
	} else {
		total = limit - zero
	}
	fmt.Println(total)

	return total
}
