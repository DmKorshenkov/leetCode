func maxProduct(nums []int) int {
    min_ := nums[0]
    max_ := nums[0]
    prod := nums[0]

    for i := 1; i < len(nums); i++ {
        if nums[i] < 0 {
            min_, max_ = max_, min_
        }

        min_ = min(min_ * nums[i], nums[i])
        max_ = max(max_ * nums[i], nums[i])
        if max_ > prod {
            prod = max_
        }
    }
    return prod
}
