func findMissingElements(nums []int) []int {
	maximum, minimum := 0, 100
	for i := 0; i < len(nums); i++ {
		maximum = max(maximum, nums[i])
		minimum = min(minimum,nums[i])
	}
	freq := make([]int, (maximum-minimum)+1)
	for i := 0; i < len(nums); i++ {
		freq[nums[i]-minimum]++
	}
	res := []int{}
	for i := range freq {
		if freq[i] == 0 {
			res = append(res, i+minimum)
		}
	}
	return res
}
