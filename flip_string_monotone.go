package main

import (
	"fmt"
	"math"
)

func main() {
	// minFlipsMonoIncr("00110")    // 1
	// minFlipsMonoIncr("010110")   // 2
	// minFlipsMonoIncr("00011000") // 2
	// minFlipsMonoIncr("01110")    // 1
	// minFlipsMonoIncr("11011")      // 2
	minFlipsMonoIncr("1111001110") // 3
}

func minFlipsMonoIncr(s string) int {
	if len(s) == 0 {
		return 0
	}

	var totalZero, totalOne int
	for _, ch := range s {
		if ch == '0' {
			totalZero++
		} else {
			totalOne++
		}
	}

	size := len(s)

	var zero, one int
	flips := make([]int, size)
	for i, ch := range s {

		if ch == '0' {
			zero++
		} else {
			one++
		}

		leftOneFlip := one
		leftZeroFlip := zero

		rightOneFlip := totalZero - zero
		rightZeroFlip := totalOne - one

		fmt.Println(leftZeroFlip, leftOneFlip, "|", rightZeroFlip, rightOneFlip)

		leftFlip := Min(leftZeroFlip, leftOneFlip)
		rightFlip := Min(rightZeroFlip, rightOneFlip)

		if leftZeroFlip < leftOneFlip {
			// fmt.Println("flip right")
			rightFlip = rightOneFlip
		}

		flips[i] = leftFlip + rightFlip
	}
	fmt.Println(flips)

	min := math.MaxInt32
	for _, v := range flips {
		if v < min {
			min = v
		}
	}
	return min
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
