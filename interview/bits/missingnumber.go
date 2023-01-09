package main

func main() {

}

func missingNumber(nums []int) int {
	n := len(nums)
	total := (n - 1) * n / 2
	for _, v := range nums {
		total -= v
	}
	return total + n
}

func missingNumber(nums []int) int {
	res := len(nums)
	for i, v := range nums {
		res ^= i
		res ^= v
	}
	return res
}
