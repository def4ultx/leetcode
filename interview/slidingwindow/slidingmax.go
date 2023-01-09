package main

import "fmt"

func main() {
	// maxSlidingWindow([]int{5, 4, 3, 2, 1, 5}, 6)
	// maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
	// maxSlidingWindow([]int{1, 2, 3, 1, 5, 7, 2, 3, 1, 5, 7, 2, 3, 1, 5, 7, 2, 3, 1, 5, 7, 2, 3, 1, 5, 7}, 4)
	for i := 100; i >= 0; i-- {
		fmt.Println(i)
	}
	maxSlidingWindow([]int{100, 99, 98, 97, 96, 95, 94, 93, 92, 91, 90, 89, 88, 87, 86, 85, 84, 83, 82, 81, 80, 79, 78, 77, 76, 75, 74, 73, 72, 71, 70, 69, 68, 67, 66, 65, 64, 63, 62, 61, 60, 59, 58, 57, 56, 55, 54, 53, 52, 51, 50, 49, 48, 47, 46, 45, 44, 43, 42, 41, 40, 39, 38, 37, 36, 35, 34, 33, 32, 31, 30, 29, 28, 27, 26, 25, 24, 23, 22, 21, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, 4)
}

var c = 0

func maxSlidingWindow(nums []int, k int) []int {
	mq := MonotonicQueue{make([]Item, 0)}
	for i := 0; i < k-1; i++ {
		c++
		mq.Append(nums[i])
	}
	fmt.Println(mq.data)
	fmt.Println("------")

	res := make([]int, 0)
	for i := k - 1; i < len(nums); i++ {
		fmt.Println("process, loop")
		mq.Append(nums[i])
		c++
		res = append(res, mq.Max())
		mq.Pop()
	}
	fmt.Println(res)
	fmt.Println(c)
	return res
}

type Item struct {
	value int
	count int
}

type MonotonicQueue struct {
	data []Item
}

func (q *MonotonicQueue) Append(x int) {
	count := 0

	for len(q.data) > 0 && x < q.data[len(q.data)-1].value {
		fmt.Println("inner loop")
		c++
		count += q.data[len(q.data)-1].count + 1
		q.data = q.data[:len(q.data)-1]
	}

	item := Item{x, count}
	q.data = append(q.data, item)
}

func (q *MonotonicQueue) Pop() {
	if len(q.data) == 0 {
		return
	}

	if q.data[0].count > 0 {
		q.data[0].count--
		return
	}
	q.data = q.data[1:]
}

func (q *MonotonicQueue) Max() int {
	if len(q.data) == 0 {
		return -1
	}
	return q.data[0].value
}
