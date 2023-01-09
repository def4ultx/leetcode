package main

import "fmt"

func main() {
	mergeSort([]int{1, 45, 22, 454, 2, 101, 39, 82, 54, 98, 4564, 78, 23, 768, 33, 1, 2, 1, 56, 55})
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2

	l := mergeSort(arr[:mid])
	r := mergeSort(arr[mid:])

	xs := merge(l, r)

	fmt.Println(xs)
	return xs
}

func merge(left, right []int) []int {
	i, j := 0, 0

	arr := make([]int, 0)
	for {
		if i >= len(left) && j >= len(right) {
			break
		} else if i >= len(left) {
			arr = append(arr, right[j])
			j++
		} else if j >= len(right) {
			arr = append(arr, left[i])
			i++
		} else if left[i] < right[j] {
			arr = append(arr, left[i])
			i++
		} else {
			arr = append(arr, right[j])
			j++
		}
	}
	return arr
}
