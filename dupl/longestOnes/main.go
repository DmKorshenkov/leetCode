func longestOnes(nums []int, k int) int {
	start := 0
	Ones := 0
	MaxLen := 0

	for end := 0; end < len(nums); end++ {
		if nums[end] == 0 {
			Ones++
		}

		if Ones > k {
			if nums[start] == 0 {
				Ones--
			}
			start++
		}

		MaxLen = max(MaxLen, (end-start)+1)
	}

	return MaxLen
}
