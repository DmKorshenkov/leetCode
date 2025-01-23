package main

func main() {

}

func plusOne(digits []int) []int {
	if digits[len(digits)-1] != 9 {
		digits[len(digits)-1]++
		return digits
	}

	var ok = true

	for in := len(digits) - 1; in > -1; in-- {
		if ok {
			digits[in]++
			if digits[in] == 10 {
				digits[in] = 0
			} else {
				ok = false
			}
		}
	}

	if ok {
		var arr = make([]int, len(digits)+1)
		arr[0] = 1
		return arr
	}
	return digits
}
