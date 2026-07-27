func truncateSentence(s string, k int) string {
	for i, countSpace := 0, 0; i < len(s); i++ {
		if s[i] == ' ' {
			countSpace++
			if countSpace == k {
				return s[:i]
			}
		}
	}
	return s
}
