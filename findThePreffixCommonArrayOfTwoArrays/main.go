func findThePrefixCommonArray(A []int, B []int) []int {  
    res := make([]int, len(A))
    arr := make([]int, len(A)+1)
 //   mpB := make(map[int]int)
    var sum int

    sum = 0
    for i := range A {
        arr[A[i]]++
        if arr[A[i]] == 2 {
            sum++
        }

        arr[B[i]]++
        if arr[B[i]] == 2 {
            sum++
        }

        res[i] = sum
    }
    return res
}



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
