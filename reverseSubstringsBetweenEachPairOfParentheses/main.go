func reverseParentheses(s string) string {
	sB := []byte(s)
	bb := bytes.Buffer{}
	//	sB := bytes.Buffer
	stack := []int{}

	for i := 0; i < len(s); i++ {

		if s[i] == byte('(') {
			stack = append(stack, i)
		} else if s[i] == byte(')') {

			for iR, jR := stack[len(stack)-1], i; iR < jR; iR, jR = iR+1, jR-1 {
				sB[iR], sB[jR] = sB[jR], sB[iR]
			}

			stack = stack[:len(stack)-1]
		}
	}

	for i := range sB {
		if sB[i] != byte('(') && sB[i] != byte(')') {
			bb.WriteByte(sB[i])
		}
	}
	return bb.String()
}
