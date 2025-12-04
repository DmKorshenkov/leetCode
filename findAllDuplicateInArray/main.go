func findDuplicates(nums []int) []int {
		res := make([]int, 0, 2)
		for i := range nums {
			in := abs(nums[i]) - 1
			if nums[in] > 0 {
				nums[in] = -nums[in]
			}
			//		if nums[i] > 0 {
			//			nums[nums[i]] = (-nums[nums[i]])
			//		} else {

//			fmt.Println(nums)
		}

//		fmt.Println(nums)
		for i := range nums {
			if nums[i] > 0 {
				res = append(res, nums[i])
			}
		}
		return res
}

func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}
