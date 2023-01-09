package main

import "fmt"

func main() {
	dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})
}

func dailyTemperatures(temperatures []int) []int {
	// create stack
	// if temp > top, set top(prev) = 1
	// then backfill until top > temp
	// then pop and push stack
	// else push to stack
	// after loop, backfill with the index with 0??, no need because there is no day with higher

	if len(temperatures) == 0 {
		return []int{}
	}

	type Temp struct {
		tempurature int
		index       int
	}

	temps := make([]Temp, len(temperatures))
	for i, v := range temperatures {
		temps[i] = Temp{v, i}
	}
	result := make([]int, len(temperatures))

	stack := []Temp{temps[0]}

	for i := 1; i < len(temps); i++ {
		fmt.Println("visit index", i, stack)
		curr := temps[i]
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			if curr.tempurature <= top.tempurature {
				break
			}
			stack = stack[:len(stack)-1]
			result[top.index] = i - top.index
		}
		stack = append(stack, temps[i])
	}

	fmt.Println(result)
	return result
}
