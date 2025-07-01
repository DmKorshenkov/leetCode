func searchMatrix(matrix [][]int, target int) bool {
    m := len(matrix)
    n := len(matrix[0])

    for u, d := 0, m -1; u <= d; {
        mid := d - (d - u) / 2

        if matrix[mid][n-1] >= target && matrix[mid][0] <= target {
            return binSearch(matrix[mid], target)
        } else if matrix[mid][n-1] < target {
            u = mid + 1
        } else if matrix[mid][n-1] > target {
            d = mid - 1
        }
    } 
    return false
}

func binSearch(nums []int, num int) bool {
	var mid int

	mid = 0
	low, hight := 0, len(nums)-1

	for low <= hight {
		mid = (low + hight) / 2
		if nums[mid] == num {
			return true
		}

		if nums[mid] < num {
			low = mid + 1
		} else {
			hight = mid - 1
		}

	}
	return false
}
