func findMiddleIndex(nums []int) int {
    var sum,sum2 int

    for i := 0; i < len(nums); i ++{
        sum2 = 0
        for j := len(nums)-1; j > i-1; j--{
            sum2 += nums[j]
        }
        if sum == sum2-nums[i]{
            return i
        }
        sum += nums[i]
    }

    return -1
}
