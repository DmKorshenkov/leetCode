func longestCommonSubsequence(text1 string, text2 string) int {

	res := make([][]int, len(text1)+1)

	for i := range res {
		res[i] = make([]int, len(text2)+1)
	}

	for i := 1; i < len(res); i++ {

		for j := 1; j < len(res[i]); j++ {
			if text1[i-1] == text2[j-1] {
				res[i][j] = res[i-1][j-1] + 1
			} else {
				res[i][j] = max(res[i-1][j], res[i][j-1])
			}
		}
	}
	return res[len(text1)][len(text2)]
}
