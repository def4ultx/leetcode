package main

import "math"

func main() {
	// reverse(321)
	reverse(-123)
}

func reverse(x int) int {
	var res int
	for x != 0 {
		rem := x % 10
		x = x / 10
		res = res*10 + rem
		if res > math.MaxInt32 || res < math.MinInt32 {
			return 0
		}
	}
	return res
}
