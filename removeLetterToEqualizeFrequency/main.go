func equalFrequency(word string) bool {
	var hash = make([]int, 26)
	var count map[int]struct{}

	for c := range word {
		hash[word[c]-97]++

	}
	fmt.Println(hash)
	for i := range hash {
		hash[i]--
		count = make(map[int]struct{})
		for j := range hash {
			count[hash[j]] = struct{}{}
		}
		if len(count) <= 2 {
			return true
		}
		hash[i]++
	}
	return false
}
