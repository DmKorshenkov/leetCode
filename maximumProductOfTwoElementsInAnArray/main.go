func maxProduct(nums []int) int {
    for i := 0; i < len(nums); i++ {
        if nums[i] > nums[0] {
            nums[0], nums[i] = nums[i], nums[0]
        }
    }
    for i := 1; i < len(nums); i++ {
        if nums[i] > nums[1] {
            nums[1], nums[i] = nums[i], nums[1]
        }
    }
    return (nums[0]-1)*(nums[1]-1)
}
