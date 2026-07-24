func longestSubarray(nums []int) int {
    maxPosibleAnd := 0

    for i := range nums {
        maxPosibleAnd = max(maxPosibleAnd, nums[i])
    }
    count := 0
    maxCount := 0
    for i := range nums {
        if nums[i] == maxPosibleAnd {
            count++
            maxCount = max(maxCount, count)
        } else {
            count = 0
        }
    }
    return maxCount
}
