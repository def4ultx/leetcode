package main

func main() {
	isInterleave("aabcc", "dbbca", "aadbbcbcac")
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s3) != len(s1)+len(s2) {
		return false
	}

	dp := make([][]bool, len(s1)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s2)+1)
	}

	dp[0][0] = true
	for i := 1; i < len(s1)+1; i++ {
		dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
	}
	for i := 1; i < len(s2)+1; i++ {
		dp[0][i] = dp[0][i-1] && s2[i-1] == s3[i-1]
	}

	for i := 1; i < len(s1)+1; i++ {
		for j := 1; j < len(s2)+1; j++ {
			a := dp[i][j-1] && s3[i+j-1] == s2[j-1]
			b := dp[i-1][j] && s3[i+j-1] == s1[i-1]
			dp[i][j] = a || b

		}
	}

	return dp[len(s1)][len(s2)]

	// dp := make([][]bool, len(s1)+1)
	// for i := range dp {
	// 	dp[i] = make([]bool, len(s2)+1)
	// }
	// dp[0][0] = true

	// for i := 0; i < len(s1)+1; i++ {
	// 	fmt.Println("loop", i)
	// 	for j := 0; j < len(s2)+1; j++ {

	// 		// if i == 0 {

	// 		// } else if j == 0 {

	// 		// } else {
	// 		// }
	// 		fmt.Println(i, j)

	// 		if i == 0 && j == 0 {
	// 			continue
	// 		} else if i == 0 {
	// 			dp[i][j] = dp[i][j-1] && s2[j-1] == s3[j-1]
	// 		} else if j == 0 {
	// 			dp[i][j] = dp[i-1][j] && s1[i-1] == s3[i-1]
	// 		} else {
	// 			// dp[i][j] = (dp[i-1][j] && s1.charAt(i-1) == s3.charAt(i+j-1)) || (dp[i][j-1] && s2.charAt(j-1) == s3.charAt(i+j-1))
	// 			// fmt.Println(i, j, "|", string(s3[i+j]))
	// 			// fmt.Println(s3[i+j-2] == s1[i-1], s3[i+j-2] == s2[j-1])
	// 			a := dp[i][j-1] && s3[i+j-1] == s2[j-1]
	// 			b := dp[i-1][j] && s3[i+j-1] == s1[i-1]
	// 			fmt.Println("curr", string(s3[i+j-1]), "-", string(s1[i-1]), dp[i-1][j], string(s2[j-1]), i, j, "|", a, b)
	// 			// c := dp[i-1][j-1] && (s3[i+j] == s1[i] || s3[i+j] == s2[j])
	// 			dp[i][j] = a || b
	// 		}

	// 	}
	// 	fmt.Println()
	// }

	// fmt.Println(dp)

	// return dp[len(s1)][len(s2)]
}
