func numOfStrings(patterns []string, word string) int {
	lnWord := len(word)
	lnPattern := len(patterns[0])

	count := 0
	for i := 0; i < len(patterns); i++ {
		lnPattern = len(patterns[i])
		for j := 0; j+lnPattern <= lnWord; j++ {
			if word[j:j+lnPattern] == patterns[i] {
				count++
				break
			}
		}
	}

	return count
}
