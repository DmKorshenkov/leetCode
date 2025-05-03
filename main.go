package main

func clearDigits2(s string) string {
	sB := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {

		if s[i] >= 48 && s[i] <= 57 {
			sB = sB[:len(sB)-1]
		} else {
			sB = append(sB, s[i])
		}
	}
	return string(sB)
}

func clearDigits(s string) string {
	//	sB := []byte{}
	sB := []byte(s)

	for i := 1; i < len(sB); i++ {
		//		fmt.Println(string(sB))
		if sB[i] <= 57 && sB[i] >= 48 {
			sB = remove(sB, i-1)
			sB = remove(sB, i-1)
			i -= 2
		}
	}
	return string(sB)
}

func remove(s []byte, index int) []byte {
	if len(s) == 1 {
		return nil
	}
	if index < 0 || index >= len(s) {
		return s
	}
	return append(s[:index], s[index+1:]...)
}
