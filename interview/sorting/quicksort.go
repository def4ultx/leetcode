package main

func main() {

}

func quickSort(arr []int) {

}

func sort(arr []int, left, right int) {
	if left < right {
		// select pivot
		// partition
		// quicksort left, quicksort right
	}
}

func partition(arr []int) int {
	pivot := arr[0]
	left := 1

	for right := len(arr) - 1; right > left; right-- {
		if arr[right] < pivot {
			arr[left], arr[right] = arr[right], arr[left]
			left++
		}
	}

	return left
}
