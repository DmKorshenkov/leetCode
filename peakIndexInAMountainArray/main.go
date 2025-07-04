func peakIndexInMountainArray(arr []int) int {
    lo, hi := 0, len(arr)-1

    for lo < hi {
        mid := lo + (hi - lo)/2

        if arr[mid] > arr[mid+1] {
            hi = mid
        } else {
            lo = mid + 1
        }
    }
    return lo
}
