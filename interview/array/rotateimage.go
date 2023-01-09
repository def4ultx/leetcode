package main

import "fmt"

func main() {
	// matrix := [][]int{
	// 	{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
	// }
	// transpose(matrix)

	matrix := [][]int{
		{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16},
	}
	transpose(matrix)
}

func rotate(matrix [][]int) {
	size, l := len(matrix), len(matrix)-1
	for i := 0; i < (size+1)/2; i++ {
		for j := 0; j < size/2; j++ {
			// fmt.Println(i, j, i == 2, i == 3)
			a, b, c, d := matrix[i][j], matrix[j][l-i], matrix[l-i][l-j], matrix[l-j][i]
			// fmt.Println(a, b, c, d)

			matrix[j][l-i] = a
			matrix[l-i][l-j] = b
			matrix[l-j][i] = c
			matrix[i][j] = d

		}
	}
	// fmt.Println(matrix)

	// n, m, l := (len(matrix)+1)/2, len(matrix)/2, len(matrix)-1
	// for i := 0; i < n; i++ {
	// 	for j := 0; j < m; j++ {
	// 		matrix[j][l-i], matrix[l-i][l-j], matrix[l-j][i], matrix[i][j] = matrix[i][j], matrix[j][l-i], matrix[l-i][l-j], matrix[l-j][i]
	// 	}
	// }
}

func transpose(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			fmt.Println(i, j)
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	fmt.Println(matrix)

	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}
	fmt.Println(matrix)
}
