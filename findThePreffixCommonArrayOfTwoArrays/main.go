func findThePrefixCommonArray(A []int, B []int) []int {  
    res := make([]int, len(A))
    mp := make(map[int]int)
 //   mpB := make(map[int]int)

    for i := 0; i < len(A); i++ {
        mp[A[i]]++
        for j := 0; j <= i; j++ {
            if _, ok := mp[B[j]]; ok {
                res[i]++
            }
        }
    }
    return res
}
