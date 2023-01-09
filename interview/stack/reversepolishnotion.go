package main

import (
	"fmt"
	"strconv"
)

func main() {
	evalRPN([]string{"4", "13", "5", "/", "+"})
}

func evalRPN(tokens []string) int {
	stack := make([]int, 0)

	for _, v := range tokens {
		if v == "+" {
			x := stack[len(stack)-2]
			y := stack[len(stack)-1]

			stack[len(stack)-2] = x + y
			stack = stack[:len(stack)-1]

		} else if v == "-" {
			x := stack[len(stack)-2]
			y := stack[len(stack)-1]

			stack[len(stack)-2] = x - y
			stack = stack[:len(stack)-1]

		} else if v == "*" {
			x := stack[len(stack)-2]
			y := stack[len(stack)-1]

			stack[len(stack)-2] = x * y
			stack = stack[:len(stack)-1]

		} else if v == "/" {
			x := stack[len(stack)-2]
			y := stack[len(stack)-1]

			fmt.Println(x, y, x/y)

			stack[len(stack)-2] = x / y
			stack = stack[:len(stack)-1]
		} else {
			num, _ := strconv.Atoi(v)
			stack = append(stack, num)
		}
	}

	fmt.Println(stack)

	result := 0
	for _, v := range stack {
		result += v
	}
	return result
}
