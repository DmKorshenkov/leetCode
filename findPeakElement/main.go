func findPeakElement(nums []int) int {
    var peak, res int
    ok := false
    res = 0
    peak = nums[0]
    for i := 0; i < len(nums)-1; i++ {
        if nums[i] > nums[i+1] {
            ok = true
            if peak < nums[i] {
                peak = nums[i]
                res = i
            }

        }
    }
    if !ok {
        return len(nums)-1
    }
    return res
}

// 07.05.26

func findPeakElement2(nums []int) int {
    l, r := 0, len(nums)-1

    for l < r {
        mid := l + (r-l)/2
        if nums[mid] < nums[mid+1] {
            l = mid + 1
        } else {
            r = mid
        }
    }
    return l
}
