func sequentialDigits(low int, high int) []int {
	digits := []int{12, 23, 34, 45, 56, 67, 78, 89, 123, 234, 345, 456, 567, 678, 789, 1234, 2345, 3456, 4567, 5678, 6789, 12345, 23456, 34567, 45678, 56789, 123456, 234567, 345678, 456789, 1234567, 2345678, 3456789, 12345678, 23456789, 123456789}
	if low > digits[len(digits)-1] || high < digits[0] {
		return nil
	}

	start, end := 0, 0
	for end < len(digits)-1 {
		if digits[start] < low {
			start++
		}
		if digits[end+1] > high {
			break
		}
		end++
	}
	if start == end {
		return []int{digits[start]}
	}
	return digits[start:end+1]
}
