package main

func main() {

}

func longestCommonSubsequence(text1 string, text2 string) int {
	distances := make([][]int, len(text1)+1)
	for i := range distances {
		distances[i] = make([]int, len(text2)+1)
	}

	for row := 1; row < len(text1)+1; row++ {
		for col := 1; col < len(text2)+1; col++ {
			if text1[row-1] == text2[col-1] {
				distances[row][col] = 1 + distances[row-1][col-1]
			} else {
				a, b := distances[row-1][col], distances[row][col-1]
				distances[row][col] = Max(a, b)
			}
		}
	}

	return distances[len(text1)][len(text2)]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
