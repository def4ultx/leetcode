package main

import (
	"fmt"
	"math"
)

func main() {
	// divide(10, 3)
	// divide(7, 3)
	// divide(7, -3)

	// divide(-1, -1)
	// divide(1, -1)
	// divide(-1, 1)
	// divide(1, 1)

	divide(-2147483648, -1)
}

func divide(dividend int, divisor int) int {
	n := 0
	p1, p2 := dividend >= 0, divisor >= 0
	dividend, divisor = Abs(dividend), Abs(divisor)
	for dividend >= divisor {
		dividend -= divisor
		n++
		if n == math.MaxInt32 || n == math.MinInt32 {
			return n
		}
	}
	fmt.Println(n, p1, p2)
	if p1 && p2 || (!p1 && !p2) {
		fmt.Println(n)
		return n
	}
	fmt.Println(-n)
	return -n
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
