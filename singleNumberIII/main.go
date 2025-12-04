func singleNumber(nums []int) []int {
		res := make([]int, 0, 2)
		hash := make(map[int]int)
		for i := 0; i < len(nums); i++ {
			hash[nums[i]]++
			if hash[nums[i]] == 2 {
				delete(hash, nums[i])
			}
		}

		for i := range hash {
			res = append(res, i)
		}

		return res
}
