func search(nums []int, target int) int {

    low, hight := 0, len(nums)-1

    for low <= hight {

        mid := (low + hight) / 2

        if nums[mid] == target {
            return mid
        }

        if nums[low] <= nums[mid] {
            if nums[low] <= target && target < nums[mid] {
                hight = mid - 1
            } else {
                low = mid + 1
            }
        } else {
            if nums[mid] < target && target <= nums[hight] {
                low = mid + 1
            } else {
                hight = mid - 1
            }
        }

    }

    return -1
}
