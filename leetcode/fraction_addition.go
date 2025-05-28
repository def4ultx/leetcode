package main

import (
	"fmt"
	"strconv"
)

func main() {
	// fractionAddition("-1/2+1/2+1/3")
	fractionAddition("-5/2+10/3+7/9")
}

func fraction(expr string) string {
	numerator := 1
	denominator := 1

	i := 0
	for {
		num := 0
		deno := 1

		negative := false
		if expr[i] == '+' || expr[i] == '-' {
			if expr[i] == '-' {
				negative = true
			}
			i++
		}

		for isDigit(expr[i]) {
			num = num*10 + int(expr[i]-'0')
			i++
		}

		i++

		for isDigit(expr[i]) {
			deno = deno*10 + int(expr[i]-'0')
			i++
		}

		if negative {
			num = -num
		}

		numerator = numerator*deno + num*denominator
		denominator = denominator * deno

		gcd := GCD(numerator, denominator)

		numerator /= gcd
		denominator /= gcd
	}
}

func isDigit(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func fractionAddition(expr string) string {
	fmt.Println(expr, len(expr))

	if expr[0] != '-' {
		expr = "+" + expr
	}
	var (
		lcm     = 2 * 3 * 2 * 2 * 5 * 7 * 3
		current int
	)
	for i, count := 0, 0; i < len(expr); i = i + count + 4 {
		fmt.Println("index", i)
		count = 0

		var (
			numerator   int
			denominator int
		)
		if expr[i+2] != '/' {
			numerator = 10
			count++
		} else {
			numerator = int(expr[i+1] - '0')
		}

		if i+count+4 < len(expr) && expr[i+count+4] != '+' && expr[i+count+4] != '-' {
			denominator = 10
			count++
		} else {
			denominator = int(expr[i+count+3] - '0')
		}

		var (
			multiplier = lcm / denominator
			number     = multiplier * numerator
		)

		fmt.Println(numerator, denominator)
		fmt.Println(expr[i+2] != '/')
		fmt.Println(i+count+4 < len(expr) && (expr[i+count+4] != '+' && expr[i+count+4] != '-'))

		if expr[i] == '+' {
			current += number
		}
		if expr[i] == '-' {
			current -= number
		}

		fmt.Println("---")
	}

	var negative bool
	if current < 0 {
		negative = true
		current = -current
	}

	div := []int{2, 2, 2, 3, 3, 5, 7}
	for i := 0; i < len(div); i++ {
		if current%div[i] == 0 {
			current = current / div[i]
			lcm = lcm / div[i]
		}
	}

	var result string
	if negative {
		result += "-"
	}

	result += strconv.Itoa(current) + "/" + strconv.Itoa(lcm)
	fmt.Println(result)
	return result
}
