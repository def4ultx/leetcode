package main

import (
	"fmt"
	"strconv"
)

func main() {
	calculate("3+2*2")
	calculate(" 3/2 ")
	calculate(" 3+5 / 2")
	calculate("22-3*5")
	calculate("1-1+1")
	calculate("0-99999")
}

func calculate(s string) int {
	if s == "" {
		return 0
	}
	var digit string
	operator := '+'

	stack := make([]int, 0)
	s += "+0"
	for _, r := range s {
		if r == ' ' {
			continue
		} else if r >= '0' && r <= '9' {
			digit += string(r)
		} else if isOperator(r) {
			if operator == '+' {
				number, _ := strconv.Atoi(digit)
				stack = append(stack, number)

				// fmt.Println("+", number)
			} else if operator == '-' {
				number, _ := strconv.Atoi(digit)
				stack = append(stack, -number)
			} else if operator == '*' {
				number, _ := strconv.Atoi(digit)
				top := stack[len(stack)-1]

				stack[len(stack)-1] = top * number

				// fmt.Println("*", number)
			} else if operator == '/' {
				number, _ := strconv.Atoi(digit)
				top := stack[len(stack)-1]

				stack[len(stack)-1] = top / number
			}
			digit = ""
			operator = r
		}
	}

	// fmt.Println(digit, string(operator))

	var result int
	for _, v := range stack {
		// fmt.Println(v)
		result += v
	}
	// fmt.Println("result")
	fmt.Println(result)
	// fmt.Println("----")

	return 0
}

func isOperator(r rune) bool {
	return r == '+' || r == '-' || r == '*' || r == '/'
}
