package main

import "fmt"

func main() {
	// numSubarrayProductLessThanK([]int{10, 5, 2, 6, 3}, 100) // 11
	// numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100) // 8

	// numSubarrayProductLessThanK([]int{10, 5, 100, 2, 6}, 100) //6
	numSubarrayProductLessThanK([]int{1, 2, 3, 4, 5}, 1) //0
	// numSubarrayProductLessThanK([]int{1, 2, 3}, 0)            // 0
}

/*
[10,5,2,6,3] k = 100

10              = 10
10,5            = 50
10,5,2          = 100
10,5,2,6        = 600
10,5,2,6,3      = 1800

5               = 5
5,2             = 10
5,2,6           = 60
5,2,6,3         = 180

2               = 2
2,6             = 12
2,6,3           = 36

6               = 6
6,3             = 18

3               = 3

10
10,5
10,5,2 (2*3/2 - 0) = 3
| 10
| 10,5
| 5
5,2
5,2,6
5,2,6,3 (3*4/2 - 1*2/2) = 5
| 5
| 5,2
| 5,2,6
| 2
| 2,6
| 6
2,6,3 = (3*4/2 - 2*3/2) = 3
| 2
| 2,6
| 2,6,3
| 6
| 6,3
| 3


=> Wrong
10,5,2,6 > 100
| 10
| 10,5
| 10,5,2
| 5
| 5,2
| 2
5,2,6
5,2,6,3 > 100
| 2
| 2,6
| 2,6,3
| 6
| 6,3
| 3
2,6,3
*/

/*
[10, 5, 2, 6] k = 100

10              = 10
10,5            = 50
10,5,2          = 100
10,5,2,6        = 600

5               = 5
5,2             = 10
5,2,6           = 60

2               = 2
2,6             = 12
6               = 6

10

5
10,5

2
5,2
10,5,2


[10,5,2,6,3]
i=0
10
count+=1

i=1
10,5
count+=2

i=2
10,5,2
left++ => 5,2
count+=3

i=3
5,2,6

i=4
5,2,6,3
left+= => 2,6,3


*/

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k == 1 {
		return 0
	}
	var (
		count             = 0
		current           = 1
		left, right, last = 0, 0, 0
	)

	for left < len(nums) && right < len(nums) {
		fmt.Println("loop", left, right, "total", count, "current", current, nums[left:right+1])
		current = current * nums[right]

		if current >= k {
			fmt.Println("fount at", current, left, right)
			count += subArrayCount(right-left) - subArrayCount(last-left) // TODO: Make sure this last-left
			last = right

			for current >= k && left < len(nums) {
				current /= nums[left]
				left++
			}
		}
		right++
	}

	fmt.Println("finish loop", "total", count, "current", current, left, right)

	if current < k {
		count += subArrayCount(right-left) - subArrayCount(last-left) // TODO: Make sure this last-left
	}

	fmt.Println("count", count)
	return count
}

func subArrayCount(n int) int {
	return n * (n + 1) / 2
}

func numSubArray(nums []int, k int) int {
	if k <= 1 {
		return 0
	}

	var (
		left = 0
		ans  = 0
		prod = 1
	)
	for right := 0; right < len(nums); right++ {
		prod *= nums[right]

		for prod >= k {
			prod /= nums[left]
			left++
		}
		ans += right - left + 1
	}
	return ans
}
