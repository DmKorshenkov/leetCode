func numberOfSteps(n int) int {
    var count int
    for n > 0 {
        if n % 2 != 0 {
            n--
        } else {
            n/=2
        }
        count++
    }
    return count
}
