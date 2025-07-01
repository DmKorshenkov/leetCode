func judgeSquareSum(c int) bool {
	left := int64(0)
	right := int64(mySqrt(c))

    fmt.Println(right)
	for left <= right {
		mid := (left * left) + (right * right)

		if mid == int64(c) {
			return true
		} else if int64(c) < mid {
			right--
		} else if int64(c) > mid {
			left++
		}
	}
	return false
}

func mySqrt(x int) int {
	for i := 0; i <= x-(x/2); i++ {
		if i*i == x {
			return i
		}
		if i*i > x {
			return i - 1
		}
	}
	return 1
}
