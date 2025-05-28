package main

import "fmt"

func main() {
	fmt.Println(isHappy(19))
	fmt.Println(isHappy(1))
	fmt.Println(isHappy(2))
}

// 1^2 + 9^2 = 82
// 8^2 + 2^2 = 68
// 6^2 + 8^2 = 100
// 1^2 + 0^2 + 0^2 = 1

func isHappy(n int) bool {
	slow, fast := n, sumOfSquare(n)
	for slow != fast {
		slow = sumOfSquare(slow)
		fast = sumOfSquare(sumOfSquare(fast))
	}
	if slow == 1 {
		return true
	}
	return false
}

func sumOfSquare(n int) int {
	sum := 0
	for n > 0 {
		x := n % 10
		sum += x * x
		n = n / 10
	}
	return sum
}
