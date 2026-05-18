func findMin(nums []int) int {
    l, r := 0, len(nums)-1

    for l <= r {
        mid := l + ( r - l ) / 2

        if nums[mid] > nums[r] {
            l = mid + 1
        } else if nums[mid] < nums[l] {
            l, r = l+1, mid
        } else {
            r = mid -1
        }
    }
 
 /*   res := nums[0]
    for start, end := 0, len(nums)-1; start <= end; start, end = start+1, end-1 {
        res = min(res, nums[start], nums[end])
    } 
    return res
*/
return nums[l]
}
