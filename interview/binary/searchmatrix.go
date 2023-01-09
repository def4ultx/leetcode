package main

func main() {

}

func searchMatrix(matrix [][]int, target int) bool {
	n, m := len(matrix), len(matrix[0])

	l, r := 0, n*m-1
	for r >= l {
		mid := l + (r-l)/2
		if matrix[mid/m][mid%m] == target {
			return true
		} else if matrix[mid/m][mid%m] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}
