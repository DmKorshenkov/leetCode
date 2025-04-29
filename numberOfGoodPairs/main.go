func numIdenticalPairs(nums []int) int {
    var res int

    res = 0
    for i := range nums {
        for j := range nums {
            if nums[i] == nums[j] && i < j {
                res++
            }
        }
    }
    return res
}



func numIdenticalPairs(nums []int) int {
    var res int

    res = 0
    mp := make(map[int]int)

    for i := range nums {
        mp[i] = nums[i]
    }

    for i := 1; i < len(nums); i++ {
        for j := 0; j < i; j++ {
            if mp[j] == nums[i] {
                res++
            }
        }
    }
    return res
}
