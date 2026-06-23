func maximumProduct(arr []int) int {
    ln := len(arr)
    for i := 1; i < ln; i++ {
        key := arr[i]

        j := i - 1
        for j >= 0 && arr[j] > key {
            arr[j+1] = arr[j]
            j--
        }
        arr[j+1] = key
    }

    return max(arr[ln-1]*arr[ln-2]*arr[ln-3], arr[0]*arr[1]*arr[ln-1])
}
