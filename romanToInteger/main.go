func romanToInt(s string) int {
	hash := map[byte]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50,
		'C': 100, 'D': 500, 'M': 1000,
	}

	res := 0
	for i := len(s) - 1; i >= 0; i-- {
		c := s[i]
		if res >= 3 && 'I' == c {
			res -= hash[c] * 2
		} else if res >= 30 && 'X' == c {
			res -= hash[c] * 2
		} else if res >= 300 && 'C' == c {
			res -= hash[c] * 2
		}
		res += hash[c]
	}
	//	fmt.Println(len(hash))
	return res
}