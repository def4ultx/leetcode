package main

import "fmt"

func main() {
	maximumSwap(2763)
}

func maximumSwap(num int) int {
	xs := make([]int, 0)
	for num > 0 {
		xs = append(xs, num%10)
		num /= 10
	}

	fmt.Println(xs)

	var (
		left, right = 0, 0
		max         = xs[0]
		maxIdx      = 0
	)
	for i := 1; i < len(xs); i++ {
		fmt.Println(i, xs[i] > max)
		if xs[i] > max {
			max = xs[i]
			maxIdx = i
		}

		if max > xs[i] {
			left, right = maxIdx, i
		}
	}

	fmt.Println(left, right)

	xs[left], xs[right] = xs[right], xs[left]

	fmt.Println(xs)
	var res int

	digit := 1
	for _, v := range xs {
		res += v * digit
		digit *= 10
	}
	fmt.Println(res)
	return res
}
