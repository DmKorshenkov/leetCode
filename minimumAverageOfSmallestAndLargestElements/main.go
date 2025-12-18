	n := len(nums) / 2
	averages := make([]float64, n)
	maxElem, minElem := 0, 0
	for step := 0; step < n; step++ {
		maxElem, minElem = 0, 0
		//		fmt.Println("first step ----   ", nums)
		for i := 0; i < (n*2)-(step*2); i++ {
			//	fmt.Println(nums[i])
			if nums[i] > nums[maxElem] {

				maxElem = i
			} else if nums[i] < nums[minElem] {
				minElem = i
			}
		}
		averages[step] = float64(float64(nums[maxElem]+nums[minElem]) / 2)
		//		fmt.Println(maxElem, "-", nums[maxElem], minElem, "-", nums[minElem])
		removeAtClear(nums, maxElem, minElem)
		//		fmt.Println(averages)
		//		fmt.Println()
	}
	res := averages[0]
	//	fmt.Println(averages)
	for i := range averages {
		res = min(res, averages[i])
	}

	return res
}

func removeAtClear(arr []int, i, j int) []int {
	if j > i {
		j--
	}
	copy(arr[i:], arr[i+1:])
	arr[len(arr)-1] = 0
	copy(arr[j:], arr[j+1:])
	arr[len(arr)-2] = 0
	return arr[:len(arr)-2]
}
