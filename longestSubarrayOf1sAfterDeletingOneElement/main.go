func longestSubarray(nums []int) int {
//	runtime.GC()
    var res int

	res = 0

	for i := range nums {

		n := 0
		step := 0
		for j := i; j < len(nums); j++ {
			if nums[j] == 0 {
				n++
                if n > 1 {
                    break
                }
                continue
			}
			step++
		}
        res = max(res, step)
	}
    if res == len(nums) {
        res -= 1
    }
	return res
}
