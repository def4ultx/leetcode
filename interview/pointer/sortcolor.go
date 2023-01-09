package main

func main() {

}

func sortColors(nums []int) {
	temp := make([]int, 3)
	for _, v := range nums {
		temp[v]++
	}

	for i := 0; i < temp[0]; i++ {
		nums[i] = 0
	}
	for i := temp[0]; i < temp[0]+temp[1]; i++ {
		nums[i] = 1
	}
	for i := temp[0] + temp[1]; i < temp[0]+temp[1]+temp[2]; i++ {
		nums[i] = 2
	}
}
