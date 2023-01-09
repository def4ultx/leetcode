package main

import "fmt"

func main() {
	// trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
	trap([]int{4, 2, 0, 3, 2, 5})
}

func trap(height []int) int {

	total := 0
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	for left < right {
		if height[left] < height[right] {
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				total += leftMax - height[left]
			}
			fmt.Println("update left")
			left++
		} else {
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				total += rightMax - height[right]
			}
			fmt.Println("update right")
			right--
		}
	}

	// max := 0
	// temp, total := 0, 0

	// for i, v := range height {
	// 	if v >= max {
	// 		total += temp
	// 		temp = 0
	// 		max = v

	// 	} else {
	// 		temp += max - v
	// 	}
	// }

	fmt.Println(total)
	return total
}

func trap2(height []int) int {
	// loop from left
	// if current >= max, set current to max and add temp rain water
	// if current < max, add (max - current) to temp rain water
	// return temp rain water

	// loop from left
	// if curr < top, add to stack

	total := 0
	stack := make([]int, 0)

	for i, v := range height {
		for len(stack) > 0 {
			if v < height[stack[len(stack)-1]] {
				break
			}

			prev := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}

			dist := i - stack[len(stack)-1] - 1
			height := Min(v, height[stack[len(stack)-1]]) - height[prev]

			total = total + (height * dist)
		}
		stack = append(stack, i)
	}

	fmt.Println(total)
	return total
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
