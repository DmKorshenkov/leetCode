func majorityElement(nums []int) []int {
    n := len(nums)
    count := n/3+1

    sort.Ints(nums)

    res := []int{}
    candidate := nums[0]
    lastElem := math.MaxInt
    for i := 0; i < n; i++ {
        if candidate == nums[i] {
            count--
        } else {
            candidate = nums[i]
            count=n/3
        }
        if count == 0 && nums[i]!=lastElem{
            res = append(res,nums[i])
            lastElem = nums[i]
            count = n/3+1
        }
    }
    return res
}
