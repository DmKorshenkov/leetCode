func sumOfGoodIntegers(n int, k int) int {
    x, sum := max(1, n-k), 0
    
    for abs(n-x) <= k {
        if (n&x) == 0 {
            sum+=x
        }
        x++
    }
    return sum
}

func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}
