func frequencySort(s string) string {
	count := make([]int, 128)
	for i := 0; i < len(s); i++ {
		
		count[s[i]]++
	}

	res := make([]byte, 0, len(s))
	for i := 0; i < len(count); i++ {
		
		freq := count[i]
		char := i
		for j := 0; j < len(count); j++ {
			if count[j] > freq {
				freq = count[j]
				char = j
			}
		}
		for freq > 0 {
			res = append(res, byte(char))
			freq--
		}
		count[char] = freq
	}
	return string(res)

}
