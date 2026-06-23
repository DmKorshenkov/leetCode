func sortByBits(arr []int) []int {
    count := make([][]int, 14)
    for i := 0; i < len(arr); i++ {
        count[firstBit(arr[i])] = append(count[firstBit(arr[i])], arr[i])
    }
    pos := 0
    for i := range count {
        sort.Ints(count[i])
        for j := range count[i] {
            arr[pos] = count[i][j]
            pos++
        }
    }
    return arr
}

func firstBit(n int) (pos int) {
    for n > 0 {
        n &= n-1
        pos++
    }
    return
}

func sortByBitsII(arr []int) []int {

    sort.Slice(arr, func(i,j int) bool {
        if firstBit(arr[j]) < firstBit(arr[i]) {
            return false
        } else if firstBit(arr[j]) == firstBit(arr[i]) && arr[j] < arr[i] {
            return false
        }
        return true
    })
    return arr
}
