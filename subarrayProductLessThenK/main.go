func numSubarrayProductLessThanK(nums []int, k int) int {
    res := 0
    for i := 0; i < len(nums); i++ {
        currMult := nums[i]
        if nums[i] < k {
            res++
        }
        for j := i+1; j < len(nums) && currMult < k; j++ {
            if currMult * nums[j] < k {
                currMult *= nums[j]
                res++
            } else {
                break
            }
        }
    }
    return res
}
