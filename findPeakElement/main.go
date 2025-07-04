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
