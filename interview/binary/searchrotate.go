package main

import "fmt"

func main() {
	// search([]int{4, 5, 6, 7, 0, 1, 2}, 0)
	// search([]int{4, 5, 6, 7, 0, 1, 2}, 3)
	// search([]int{4, 4, 4, 7, 0, 1, 2}, 4)
	// search([]int{4, 4, 4, 7, 0, 1, 2}, 5)
	// search([]int{1, 3}, 3)
	// search([]int{3, 5, 1}, 3)
	search([]int{5, 1, 3}, 3)
	search([]int{5, 1, 3}, 5)
	// search([]int{4, 5, 6, 7, 8, 1, 2, 3}, 4)
	// search([]int{4, 5, 6, 7, 8, 1, 2, 3}, 8)
}

func search(nums []int, target int) int {

	if len(nums) == 0 {
		return -1
	}

	l, r := 0, len(nums)-1
	for r >= l {
		mid := l + (r-l)/2
		fmt.Println("visit", l, r, mid, "|")
		if nums[mid] == target {
			fmt.Println("found", mid)
			return mid
		}

		if nums[l] <= nums[mid] {
			fmt.Println("visit left")
			if target < nums[mid] && target >= nums[l] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			fmt.Println("visit right")
			if target > nums[mid] && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	fmt.Println("not found")
	return -1

	// for r >= l {
	// 	mid := l + (r-l)/2

	// 	fmt.Println("visit", l, r, mid, "|", nums[l], nums[r], nums[mid])
	// 	if nums[mid] == target {
	// 		fmt.Println("found", mid)
	// 		return mid
	// 	}

	// 	if nums[l] < nums[r] && nums[mid] < nums[r] {
	// 		fmt.Println("sorted")
	// 		// sorted
	// 		if nums[mid] > target {
	// 			r = mid - 1
	// 		} else {
	// 			l = mid + 1
	// 		}
	// 		continue
	// 	}

	// 	if nums[mid] < nums[r] && target > nums[mid] && target <= nums[r] {
	// 		fmt.Println("condition")
	// 		l = mid + 1
	// 		continue
	// 	}
	// 	if target > nums[mid] && nums[mid] > nums[l] {
	// 		fmt.Println("condition")
	// 		l = mid + 1
	// 		continue
	// 	}
	// 	if nums[r] < nums[mid] && target <= nums[r] {
	// 		fmt.Println("why")
	// 		l = mid + 1
	// 	} else {
	// 		fmt.Println("oh")
	// 		r = mid - 1
	// 	}
	// 	fmt.Println("finish visit", l, r, mid)

	// }

	// fmt.Println(result)

	// return result
}
