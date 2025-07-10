func rob(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    sum1 := nums[0]
    sum2 := max(nums[0],nums[1])

    for i := 2; i < len(nums); i++ {
        sum2, sum1 = max(sum1+nums[i], sum2), sum2
    }


    return sum2
}
