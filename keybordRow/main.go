func findWords(words []string) []string {
	keys := "qwertyuiopasdfghjklzxcvbnm"

	firstMp := make(map[byte]int)

	for i := 0; i < 10; i++ {
		firstMp[keys[i]] = 1
	}

	for i := 10; i < 19; i++ {
		firstMp[keys[i]] = 2
	}

	for i := 19; i < 26; i++ {
		firstMp[keys[i]] = 3
	}

	res := make([]string, 0, len(words))
loop:
	for i := 0; i < len(words); i++ {
		tmp := strings.ToLower(words[i])
		check := firstMp[tmp[0]]

		for j := 1; j < len(words[i]); j++ {
			if check != firstMp[tmp[j]] {
				continue loop
			}
		}
		res = append(res, words[i])
	}

	return res
}
