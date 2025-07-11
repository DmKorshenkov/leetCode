func singleNonDuplicate(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    lo, hi := 0, len(nums)-1

    for lo < hi {
        mid := lo + (hi - lo)/2
        if mid%2 == 1 {
            mid--
        }

        if nums[mid] == nums[mid+1] {
            lo = mid + 2
        } else {
            hi = mid
        }
    }
    return nums[lo]
}
