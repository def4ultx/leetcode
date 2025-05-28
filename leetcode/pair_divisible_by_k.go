package main

import "fmt"

func main() {
	// (1,9),(2,8),(3,7),(4,6) and (5,10)
	// canArrange([]int{1, 2, 3, 4, 5, 10, 6, 7, 8, 9}, 5)
	// canArrange([]int{1, 2, 3, 4, 5, 10, 6, 7, 8, 9}, 7)

	// fmt.Println(-21 % 5)
	// fmt.Println((-1 + 3) % 3)
	// fmt.Println(-11 % 7)
	// canArrange([]int{-1, 1, -2, 2, -3, 3, -4, 4}, 3)

	canArrange([]int{-1, 1}, 5)
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func canArrange(arr []int, k int) bool {

	// neg := make([]int, 0)
	// for _, v := range arr {
	// 	if v < 0 {
	// 		neg = append(neg, v)
	// 	}
	// }

	// big := make([]int, 0)
	// for _, v := range arr {
	// 	if v > k {
	// 		big = append(big, v)
	// 	}
	// }

	complements := make(map[int]int)
	// for _, v := range big {
	// 	complements[v%k]++
	// }
	for _, v := range arr {
		complements[mod(v, k)]++
	}

	fmt.Println(complements)

	// for _, v := range arr {

	// }

	// visit := make(map[int]bool)

	// result := 0
	for n, v := range complements {
		c := k - n
		// if visit[complement] {
		// 	continue
		// }

		// fmt.Println("n,v", n, v, m[complement], n == 0 && v%2 != 0)

		if n == 0 && v%2 != 0 {
			return false
		} else if v != complements[c] {
			return false
		}
	}

	fmt.Println("result=true")

	return true
}
