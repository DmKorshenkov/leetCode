func maximumProduct(nums []int) int {
	min1, min2 := 1001, 1001
	max1, max2, max3 := -1001, -1001, -1001
	for _, num := range nums {
		if num < min1 {
			min1, min2 = num, min1
		} else if num < min2 {
			min2 = num
		}

		if num > max1 {
			max1, max2, max3 = num, max1, max2
		} else if num > max2 {
			max2, max3 = num, max2
		} else if num > max3 {
			max3 = num
		}
	}

	return max(min1*min2*max1, max1*max2*max3)
}
