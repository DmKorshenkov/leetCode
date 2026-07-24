func countMaxOrSubsets(nums []int) int {
    maxOr := 0
    count := 0
    for i := range nums {
        maxOr |= nums[i]
    }

    var back func(in int, or int)

    back = func (in int, or int) {
        if in == len(nums) {
            if or == maxOr {
                count++
            }
            return
        }

        back(in+1, or)
        back(in+1, or | nums[in])
    }

    back(0,0)
    return count
}
