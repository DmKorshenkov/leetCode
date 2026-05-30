func minElement(nums []int) int {
    res := nums[0]
    N := 0
    for _, n := range nums {
        N = 0
        for n > 0 {
            N += n % 10
            n /= 10
        }

        if N < res {
            res = N
        }
    }
    return res
}
