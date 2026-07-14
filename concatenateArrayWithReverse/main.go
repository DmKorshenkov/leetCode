func concatWithReverse(nums []int) []int {
    n := len(nums)
    newArray := make([]int, 0, n*2)
    newArray = append(newArray, nums...)
    for n > 0 {
        n--
        newArray = append(newArray, nums[n])
    }
    return newArray
}
