func rotate(nums []int, k int) {
    for ; k > 0; k-- {
        last := nums[len(nums)-1]
        copy(nums[1:], nums[:len(nums)-1])
        nums[0] = last
    }
    
///    return nums
}
