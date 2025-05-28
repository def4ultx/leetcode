package main

func main() {

}

func searchRange(nums []int, target int, leftSearch bool) []int {

	l := subSearchRange(nums, target, true)
	r := subSearchRange(nums, target, false)
	return []int{l, r}
}

func subSearchRange(nums []int, target int, leftSearch bool) int {
	l, r := 0, len(nums)-1

	mid := -1
	for l <= r {
		m := l + (r-l)/2

		if nums[m] == target {

			mid = m
			if leftSearch {
				r = m - 1
			} else {
				l = m + 1
			}
		} else if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return mid
}

func searchRange(nums []int, target int) []int {
	l, r := 0, len(nums)-1

	for l <= r {
		m := l + (r-l)/2

		if nums[m] == target {
			l, r = m, m
			for i := m - 1; i >= 0; i-- {
				if nums[i] == target {
					l = i
				}
			}
			for i := m + 1; i < len(nums); i++ {
				if nums[i] == target {
					r = i
				}
			}

			return []int{l, r}
		}

		if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return []int{-1, -1}
}
