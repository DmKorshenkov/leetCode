func kLengthApart(nums []int, k int) bool {

	startIn := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			if startIn != -1 {
				check := (i - startIn) - 1
				if check < k {
					return false
				}
			}
			startIn = i
		}
	}

	return true
}
