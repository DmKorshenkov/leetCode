func searchRange(nums []int, target int) []int {
	low, high := 0, len(nums)-1
	start := -1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			start = mid
		}

		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	end := -1
	low, high = 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			end = mid
		}
		if nums[mid] <= target {
			low = mid + 1

		} else {
			high = mid - 1
		}
	}

	return []int{start, end}
}
