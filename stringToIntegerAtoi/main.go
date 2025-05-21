func myAtoi(s string) int {
    var i int
	var neg int

	neg = 1
	i = 0
	//	s := []byte(s)

	for i < len(s) && s[i] == byte(' ') {
		i++
	}

	if i < len(s) && s[i] == byte('-') {
		neg = -1
		i++
	} else if i < len(s) && s[i] == byte('+') {
        i++
    }

	for i < len(s) && s[i] == byte('0') {
		i++
	}

	res := 0
	for i < len(s) {
		if !(s[i] >= byte('0') && s[i] <= byte('9')) {
			return res * neg
		}
		res *= 10
		res += int(s[i] - 48)

		if res >= 2147483647 && neg == 1 {
			return 2147483647
		}

		if res >= 2147483648 && neg == -1 {
			return -2147483648
		}
		i++
	}

	return res * neg
}
