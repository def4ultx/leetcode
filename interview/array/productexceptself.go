package main

func main() {

}

func productExceptSelf(nums []int) []int {

	size := len(nums)
	result := make([]int, size)
	for i := range result {
		result[i] = 1
	}

	left, right := 1, 1
	for i := 0; i < size; i++ {
		result[i] *= left
		left = left * nums[i]
	}
	for i := size - 1; i >= 0; i-- {
		result[i] *= right
		right = right * nums[i]
	}

	return result
}
