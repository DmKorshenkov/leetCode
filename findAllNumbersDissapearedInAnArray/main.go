func findDisappearedNumbers(nums []int) []int {
   
   for _, n := range nums {
        indx := abs(n) - 1
        if nums[indx] > 0 {
            nums[indx] = -nums[indx]
        }
    }

    var res = make([]int, 0, len(nums))
//    res := []int{}
    for i, n := range nums {
        if n > 0 {
           res = append(res, i+1) 
        }
    }
    return res
}

func abs(a int) int {
    if a > 0 {
        return a
    }

    return -a
}

func findDisappearedNumbers(nums []int) []int {
    sort.Ints(nums)
    n := len(nums)
    arr := make([]int, 0, n)

    for i := 1; i < n+1; i++ {
        if binSearch(nums, i) == - 1{
            arr = append(arr, i)
        }
    }

    return arr
}

func binSearch(nums []int, num int) int {
	var mid int

	mid = 0
	low, hight := 0, len(nums)-1

	for low <= hight {
		mid = (low + hight) / 2
		if nums[mid] == num {
			return num
		}

		if nums[mid] < num {
			low = mid + 1
		} else {
			hight = mid - 1
		}

	}
	return -1
}
