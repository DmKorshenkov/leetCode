func getFinalState(nums []int, k int, multiplier int) []int {
    n := len(nums)
    for i := 0; i < k; i++ {
        minimum := 0
        for j := n-1; j >= 0; j-- {
            if nums[j] <= nums[minimum] {
                minimum = j
            }
        }
        nums[minimum] *= multiplier
    }
    return nums
}
