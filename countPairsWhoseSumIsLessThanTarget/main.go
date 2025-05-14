func countPairs(nums []int, target int) int {
    var res int

    res = 0
//    sort.Ints(nums)
    for i := 0; i < len(nums); i++ {
        for j := len(nums)-1; j > i; j-- {
            if nums[i] + nums[j] < target {
                res++
            }
        }
    }
    return res
}
