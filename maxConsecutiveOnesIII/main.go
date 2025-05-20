func longestOnes(nums []int, k int) int {
	var res int

	res = 0

	for i := range nums {

		n := 0
		step := 0

		for j := i; j < len(nums); j++ {
			if nums[j] == 0 {
				n++
			}
			if n > k {
				break
			}
			step++
		}

		if step > res {
			res = step
		}

	}
//	fmt.Println(res)
	return res
}
