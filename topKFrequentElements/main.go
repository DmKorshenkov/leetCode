func topKFrequent(nums []int, k int) []int {
	min := nums[0]
	max := nums[0]
	// {1,2,2,1,3,-1}
	for i := range nums {
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
	}

	count := make([]int, (max+1)+(-min))
	// {0{1},1{2},2{2},3{1}}
	for i := range nums {
		count[nums[i]-min]++
	}
	//	fmt.Println(count)

	bucket := make([][]int, len(nums))
	for i := range nums {
		indexNum := nums[i] - min
		countNum := count[indexNum]
		if countNum != 0 {
			bucket[countNum-1] = append(bucket[countNum-1], nums[i])
		}
		count[indexNum] = 0
	}
	//	fmt.Println(bucket)
	result := make([]int, 0, k)
	for i := len(nums) - 1; k > 0; i-- {
		if bucket[i] != nil {
			result = append(result, bucket[i]...)
			k -= len(bucket[i])
		}
	}

	return result
}
