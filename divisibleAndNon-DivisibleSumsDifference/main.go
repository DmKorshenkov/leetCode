func differenceOfSums(n int, m int) int {
    var sum int
    for n > 0 {
        if n % m != 0 {
            sum += n
        } else {
            sum -= n
        }
        n--
    }
    return sum
}
