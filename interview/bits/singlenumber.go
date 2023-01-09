package main

func main() {

}

func singleNumber(nums []int) int {
	x := 0
	for i := range nums {
		x = x ^ nums[i]
	}
	return x
}
