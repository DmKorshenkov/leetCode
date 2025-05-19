func rearrangeArray(nums []int) []int {
    res := make([]int, len(nums))
    positive := 0
    negative := 1
    for i := 0; i < len(nums); i++ {
        if nums[i] < 0 {
            res[negative] = nums[i]
            negative += 2
        } else {
            res[positive] = nums[i]
            positive += 2
        }
    }
    return res
}
