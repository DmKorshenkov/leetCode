func alternatingSum(nums []int) int {
    var sum int 

    for i := range nums {
        if i % 2 == 0 {
            sum+=nums[i]
        } else {
            sum-=nums[i]
        }
    }
    return sum
}
