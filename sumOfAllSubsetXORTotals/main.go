func subsetXORSum(nums []int) int {
	res := 0

	var back func(in int, xor int)

	back = func(in int, xor int) {
		if in == len(nums) {
			res += xor
			return
		}

		back(in+1, xor)
		back(in+1, xor ^ nums[in])
	}
	back(0,0)
	return res
}
