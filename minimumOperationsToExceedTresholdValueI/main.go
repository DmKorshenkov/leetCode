func minOperations(nums []int, k int) int {
    var count int
    for i := range nums {
        if nums[i] < k {
            count++
        }
    }
    return count
}
