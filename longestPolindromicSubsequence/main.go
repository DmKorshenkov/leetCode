func longestPalindromeSubseq(s string) int {
	str1 := reverse(s)

	dp := make([][]int, len(str1)+1)

	for i := range dp {
		dp[i] = make([]int, len(dp)+1)
	}

	for i := 1; i < len(str1)+1; i++ {
		for j := 1; j < len(str1)+1; j++ {
			if s[i-1] == str1[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(dp)-1][len(dp)-1]
}

func reverse(str string) string {
    bb := bytes.Buffer{}
    bb.WriteString(str)

    for i, j := 0, bb.Len()-1; i < j; i,j =i+1,j-1 {
        bb.Bytes()[i],bb.Bytes()[j]=bb.Bytes()[j],bb.Bytes()[i]
    }
    return bb.String()
}
