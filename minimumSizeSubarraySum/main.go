func minSubArrayLen(target int, nums []int) int {
    ln := len(nums)
    res := ln + 1
    start := 0
    
    windowSum := 0
    for end := 0; end < ln; end++{
        windowSum += nums[end]

        for windowSum >= target {

            if (end - start) + 1 < res {
                res = (end - start) + 1
            }

            windowSum -= nums[start]
            start++
        }
    }
    if res == ln + 1 {
        return 0
    }
    return res 
}
