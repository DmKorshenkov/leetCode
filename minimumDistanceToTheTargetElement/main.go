func getMinDistance(nums []int, target int, start int) int {

    left, right := start, start

    for {
        if left >= 0 && nums[left] == target {
            return start - left
        }

        if right < len(nums) && nums[right] == target {
            return right - start
        }
        left--
        right++
    }

    return 0
}
