func findMaxAverage(nums []int, k int) float64 {
	if len(nums) == 1 {
		return float64(nums[0])
	}
	res := float64(sum(nums[:k])) / float64(k)
	for i := 1; i+k <= len(nums); i++ {
		if average := float64(sum(nums[i:i+k])) / float64(k); average > res {
			res = average
		}
	}
	return res
}

func sum(nums []int) (res int) {
	for i := range nums {
		res += nums[i]
	}
	return res
}

