func subarraySum(nums []int) int {
    sum := 0
    for i := 0; i < len(nums); i++ {
        for start := max(0, i - nums[i]); start <= i; start++ {
            sum+=nums[start]
        }
    }
    return sum
}
