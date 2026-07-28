func subtractProductAndSum(n int) int {
    var sum, prod int
    sum = n %10
    prod = sum
    n/=10
    for n > 0 {
        sum += n%10
        prod *= n%10
        n/=10
    }
    return prod - sum
}
