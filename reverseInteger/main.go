func reverse(x int) int {
	//--2147483648
	//2147483647
	//
    maxInt32 := 2147483647
	minInt32 := -2147483648

	var ok bool
	if x < 0 {
		x = -x
		ok = true
	}

	res := 0
	for x > 0 {
		digit := x % 10

        if res > maxInt32/10 || (res == maxInt32/10 && digit > 7) {
		    return 0
		}
		if res < minInt32/10 || (res == minInt32/10 && digit < -8) {
		    return 0
		}

        res = res*10 + digit
		x /= 10
	}

	if ok {
		return -res
	}
	return res
}
