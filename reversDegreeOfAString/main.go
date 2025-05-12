func reverseDegree(s string) int {
    var res int

    res = 0

    for i := 0; i < len(s); i++ {
        res += Abs(int(byte(s[i])) - 123) * (i+1)
    }

    return res
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
