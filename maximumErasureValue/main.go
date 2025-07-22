func maximumUniqueSubarray(nums []int) int {
    seen := [10001]bool{}
    sumMax := 0
    currSum := 0
    
    left := 0

    for right := 0; right < len(nums); right++ {
        for seen[nums[right]] {
            currSum -= nums[left]
            seen[nums[left]] = false
            left++
        }

        currSum += nums[right]
        seen[nums[right]] = true
        sumMax = max(sumMax, currSum)
    }

    return sumMax
}
