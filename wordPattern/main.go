func wordPattern(pattern string, s string) bool {
	sl := strings.Split(s, " ")
	if len(sl) != len(pattern) {
		return false
	}
	hash := make(map[string]int)
	hashPattern := make(map[byte]int)

	for i := 0; i < len(sl); i++ {
		if _, ok := hash[sl[i]]; !ok {
			hash[sl[i]] = len(hash)
		}
		if _, ok := hashPattern[pattern[i]]; !ok {
			hashPattern[pattern[i]] = len(hashPattern)
		}

		if hash[sl[i]] != hashPattern[pattern[i]] {
			return false
		}
	}
	return true
}
