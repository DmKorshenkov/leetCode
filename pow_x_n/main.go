func myPow(x float64, n int) float64 {
    N := int64(n)
    if N < 0 {
        x = 1 / x
        N = -N
    }

    result := 1.0
    for N > 0 {
        if N % 2 == 1 {
            result *= x
        }

        x *= x
        N /= 2
    }
    return result
}
