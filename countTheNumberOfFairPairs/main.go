func countFairPairs(nums []int, lower int, upper int) int64 {
	var res, l, u int64
	//   var sum int

	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		l = lower_bound(nums, i+1, len(nums)-1, lower-nums[i])
		u = lower_bound(nums, i+1, len(nums)-1, upper-nums[i]+1)
		res += u - l
	}

	return res
}

func lower_bound(nums []int, left, right, num int) int64 {
	//  var mid int
	for left <= right {
		var mid = left + ((right - left) / 2)

		if nums[mid] >= num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return int64(left)
}