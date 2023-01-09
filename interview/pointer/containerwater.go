package main

func main() {

}

func maxArea(height []int) int {
	l, r := 0, len(height)-1

	max := 0
	for l < r {
		h := Min(height[l], height[r])
		area := h * (r - l)
		max = Max(max, area)

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return max
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
