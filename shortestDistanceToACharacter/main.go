	answer := make([]int, len(s))

	for i := 0; i < len(answer); i++ {
		for j := range s {
			if s[j] == c {
				answer[i] = abc(i - j)
			}
		}
	}
//	fmt.Println(answer)
	for i := len(s) - 1; i >= 0; i-- {
		for j := len(s) - 1; j >= 0; j-- {
			if s[j] == c {
				answer[i] = min(answer[i], abc(i-j))
			}
		}
	}
	return answer
