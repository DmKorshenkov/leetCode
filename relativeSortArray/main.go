func relativeSortArray(arr1 []int, arr2 []int) []int {
    l := 0
    for i := 0; i < len(arr2); i++ {

        for j := 0; j < len(arr1); j++ {

            if arr2[i] == arr1[j] {
                arr1[l], arr1[j] = arr1[j],arr1[l]
                l++
            }
        }
    }
    sort.Ints(arr1[l:])
    return arr1
}
