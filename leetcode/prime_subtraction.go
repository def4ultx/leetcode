package main

import "fmt"

func main() {
	// fmt.Println(primeSubOperation([]int{4, 9, 6, 10}))
	// fmt.Println(primeSubOperation([]int{6, 8, 11, 12}))
	// fmt.Println(primeSubOperation([]int{5, 8, 3}))
	// fmt.Println(primeSubOperation([]int{2, 2}))
	// fmt.Println(primeSubOperation([]int{15, 15, 10}))
	fmt.Println(primeSubOperation([]int{17, 2}))
}

func primeSubOperation(nums []int) bool {
	if len(nums) < 2 {
		return true
	}

	var max int
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	primes := sieveOfEratosthenes(max)
	fmt.Println(primes)

	if len(primes) == 0 {
		return false
	}

	for i := len(nums) - 2; i >= 0; i-- {
		diff := nums[i+1] - nums[i]
		if diff > 0 {
			continue
		}
		diff = Abs(diff) + 1

		fmt.Println("loop------", i, diff, nums[i], nums[i+1])

		fmt.Println("binary-------")
		l, r, idx := 0, len(primes)-1, -1
		for l <= r {
			idx = l + (r-l)/2

			fmt.Println(diff, primes[idx])

			if diff > primes[idx] {
				l = idx + 1
			} else if diff < primes[idx] {
				r = idx - 1
			} else {
				break
			}
		}
		if primes[idx] < diff && idx < len(primes)-1 {
			idx++
		}
		// fmt.Println("idx", idx, "pick prime", primes[idx])
		fmt.Println("idx", idx)

		nums[i] = nums[i] - primes[idx]
		if nums[i] <= 0 {
			return false
		}
	}

	fmt.Println(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] >= nums[i+1] {
			return false
		}
	}
	return true
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func sieveOfEratosthenes(N int) (primes []int) {
	b := make([]bool, N)
	for i := 2; i < N; i++ {
		if b[i] == true {
			continue
		}
		primes = append(primes, i)
		for k := i * i; k < N; k += i {
			b[k] = true
		}
	}
	return
}
