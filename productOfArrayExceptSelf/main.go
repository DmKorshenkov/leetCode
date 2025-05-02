func productExceptSelf(nums []int) []int {
	//	fmt.Println(nums)
	arr := make([]int, len(nums))
	arr[0] = 1

	for i := 1; i < len(nums); i++ {
		arr[i] = nums[i-1] * arr[i-1]
	}

	arr2 := make([]int, len(nums))
	arr2[len(arr2)-1] = 1
	for i := len(nums) - 2; i > -1; i-- {
		arr2[i] = nums[i+1] * arr2[i+1]
	}

	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * arr2[i]
	}
	//	fmt.Println(arr)
	return arr
}

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	for i := range res {
		res[i] = 1
	}
	var current int
	var current1 int

	current = 1
	current1 = 1

	for i := range nums {
		j := (len(nums) - 1) - i
		res[i] *= current
		current *= nums[i]

		res[j] *= current1
		current1 *= nums[j]
	}

	return res
}
