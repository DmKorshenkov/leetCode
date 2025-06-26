func minDistance(word1 string, word2 string) int {
    
	m := len(word1)
	n := len(word2)

	res := make([][]int, m+1)

	for i := range res {
		res[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {

		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				res[i][j] = res[i-1][j-1] + 1
			} else {
				res[i][j] = max(res[i-1][j], res[i][j-1])
			}
		}
	}
	return m+n - (res[m][n] * 2)
}
