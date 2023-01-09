package main

import "fmt"

func main() {
	climbStairs(2)
	climbStairs(3)
}

func climbStairs(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	fmt.Println(a, b)
	return b
}
