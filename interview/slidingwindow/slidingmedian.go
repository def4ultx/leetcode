package main

import (
	"container/list"
	"fmt"
	"sort"
)

func main() {
	medianSlidingWindowTwoHeap([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
	// [1.00000,-1.00000,-1.00000,3.00000,5.00000,6.00000]
}

// Sol 1. - maintain sorted list
// Sol 2. - using binary serach to remove last item
// Sol 3. - Using 2 heap
// Sol 4. - Using SortedList structure

func medianSlidingWindowTwoHeap(nums []int, k int) []float64 {
}

// func medianSlidingWindowBinarySearch(nums []int, k int) []float64 {
// 	temp := make([]int, k)
// 	copy(temp, nums[:k])
// 	sort.Ints(temp)

// 	res := []float64{findMedian(temp)}
// 	for i := k; i < len(nums); i++ {
// 		fmt.Println("search", nums[i-k], temp)
// 		idx := sort.Search(k, func(j int) bool {
// 			return temp[j] >= nums[i-k]
// 		})
// 		temp[idx] = nums[i]

// 		m := findMedian(temp)
// 		res = append(res, m)
// 	}
// 	fmt.Println(res)
// 	return res
// }

// func findMedian(nums []int) float64 {
// 	n := len(nums)
// 	if n%2 != 0 {
// 		return float64(nums[n/2])
// 	} else {
// 		a := float64(nums[n/2])
// 		b := float64(nums[(n/2)-1])
// 		return (a + b) / 2
// 	}
// }

func medianSlidingWindow(nums []int, k int) []float64 {

	temp := make([]int, k)
	copy(temp, nums[:k])
	sort.Ints(temp)

	ll := list.New()
	for _, v := range temp {
		ll.PushBack(v)
	}

	res := make([]float64, 0)
	res = append(res, median(ll, k))
	// Printlist(ll)

	for i := k; i < len(nums); i++ {
		fmt.Println("remove", nums[i-k], "insert", nums[i])
		Printlist(ll)
		remove(ll, nums[i-k])
		Printlist(ll)
		insert(ll, nums[i])
		Printlist(ll)
		fmt.Println("----")

		res = append(res, median(ll, k))
	}

	fmt.Println(res)
	return res
}

func median(ll *list.List, k int) float64 {
	l := ll.Front()
	for i := 0; i < k/2; i++ {
		l = l.Next()
	}

	if k%2 != 0 {
		return float64(l.Value.(int))
	} else {
		p := l.Prev()
		a := float64(p.Value.(int))
		b := float64(l.Value.(int))
		return (a + b) / 2
	}
}

func remove(ll *list.List, x int) {
	l := ll.Front()
	for l != nil {
		fmt.Println("remove", l, x, l.Value.(int) == x)
		if l.Value.(int) == x {
			ll.Remove(l)
			break
		}
		l = l.Next()
	}
}

func insert(ll *list.List, x int) {
	l := ll.Front()
	for l != nil {
		fmt.Println("insert", l)
		if l.Value.(int) >= x {
			break
		}
		l = l.Next()
	}
	if l == nil {
		ll.PushBack(x)
	} else {
		ll.InsertBefore(x, l)
	}
}

func Printlist(ll *list.List) {
	for e := ll.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value.(int), " ")
	}
	fmt.Println()
}
