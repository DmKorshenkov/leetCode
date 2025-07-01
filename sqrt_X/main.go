func mySqrt(x int) int {
    for i := 0; i <= x - (x/2); i++ {
        if i*i == x {
            return i
        }
        if i*i > x {
            return i-1
        }
    }
    return 1
}
