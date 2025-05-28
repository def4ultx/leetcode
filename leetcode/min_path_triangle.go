package main

func minimumTotal(triangle [][]int) int {

	n := len(triangle)
	for i := 1; i < len(triangle); i++ {

		l := len(triangle[i]) - 1
		triangle[i][0] = triangle[i][0] + triangle[i-1][0]
		triangle[i][l] = triangle[i][l] + triangle[i-1][l-1]

		for j := 1; j < l; j++ {
			triangle[i][j] = triangle[i][j] + Min(triangle[i-1][j], triangle[i-1][j-1])
		}
	}

	min := triangle[n-1][0]
	for _, v := range triangle[n-1] {
		min = Min(min, v)
	}
	return min
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
2,0 = [1,0]
2,1 = [1,0] [1,1]
2,2 = [1,1]

3,0 = [2,0]
3,1 = [2,0] [2,1]
3,2 = [2,1] [2,2]
3,3 = [2,2]
*/
