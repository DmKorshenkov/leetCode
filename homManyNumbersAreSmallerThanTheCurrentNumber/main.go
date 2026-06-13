func smallerNumbersThanCurrent(nums []int) (res []int) {
    var count [101]int
    res = make([]int, len(nums))
    for i := range nums {
        count[nums[i]]++
    }

    for i := 1; i < len(count); i++ {
        count[i] += count[i-1]
    }

    for i := range nums {
        if nums[i] > 0 {
            res[i] = count[nums[i] - 1]
        } else {
            res[i] = 0
        }
    }
    return
}
