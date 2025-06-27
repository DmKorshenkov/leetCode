func minimumDeleteSum(s1 string, s2 string) int {
    m := len(s1)
	n := len(s2)

	dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }
	var sumA, sumB int
    for i := range s1 {
        sumA += int(s1[i])
    }
    for i := range s2 {
        sumB += int(s2[i])
    }
    
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {

			if s1[i-1] == s2[j-1] {

				dp[i][j] = dp[i-1][j-1] + int(s1[i-1])
			} else {

				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return (sumA + sumB) - 2*int(dp[m][n])
}
