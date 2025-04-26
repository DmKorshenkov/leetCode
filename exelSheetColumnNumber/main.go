func titleToNumber(columnTitle string) int {
	var res int

    res = 0
	for i := 0; i < len(columnTitle); i++ {
		res *= 26
		res += int(columnTitle[i] - 64)
	}

	return res
}
