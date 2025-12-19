func distinctAverages(nums []int) int {
	hash := make(map[float64]int)
	sort.Ints(nums)
	fmt.Println(nums)
	for start, end := 0, len(nums)-1; start < end; start, end = start+1, end-1 {
		hash[float64(float64((nums[start]+nums[end]))/2)]++

	}
	return len(hash)
}
