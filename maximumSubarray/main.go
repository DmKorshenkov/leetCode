func maxSubArray(nums []int) int {
    sum, maxSum  := 0, math.MinInt
    for i := 0; i < len(nums); i++ {
        sum = max(sum+nums[i], nums[i])
        maxSum = max(maxSum, sum)
    }

    return maxSum
}
