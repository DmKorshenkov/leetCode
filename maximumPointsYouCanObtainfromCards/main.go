	totalPoints := 0
	n := len(cardPoints)
	winPoints := 0
	for i := range cardPoints {
		totalPoints += cardPoints[i]
	}
//	sumAndMultiply(k)
//	fmt.Println("total - ", totalPoints, "\u2714\u2714\u2714")
	for i := 0; i < n-k; i++ {
		winPoints += cardPoints[i]
	}
	//	fmt.Println("win - ", winPoints)

	//
	//
	out := totalPoints - winPoints
	//out := max(totalPoints-winPoints, winPoints)
	for end, start := n-k, 0; end < len(cardPoints); end, start = end+1, start+1 {
		winPoints += cardPoints[end]
		winPoints -= cardPoints[start]
		fmt.Println((totalPoints - winPoints))
		out = max(out, (totalPoints - winPoints))
	}
	return out
