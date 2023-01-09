package main

import "fmt"

func main() {
	insertionSort([]int{1, 45, 22, 454, 2, 101, 39, 82, 54, 98, 4564, 78, 23, 768, 33, 1, 2, 1, 56, 55})
}

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {

		j := i - 1

		for j >= 0 && arr[j] > arr[j+1] {
			arr[j], arr[j+1] = arr[j+1], arr[j]
			j--
		}
	}
	fmt.Println(arr)
}
