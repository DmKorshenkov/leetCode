func findDuplicates(nums []int) []int {		
	res := make([]int, 0, 2)
		for i := range nums {
			tmp := abs(nums[i])

			if nums[tmp-1] < 0 {
				res = append(res, tmp)
			} else {
				nums[tmp-1] *= -1
			}
//			fmt.Println(nums, res)
		}

//		fmt.Println(nums, res)
		return res
}

func abs(n int) int {

	if n < 0 {
		return -n
	}
	return n
}
