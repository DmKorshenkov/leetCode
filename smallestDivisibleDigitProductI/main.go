func smallestNumber(n int, t int) int {   
    if n <= t {
        return t
    }

    for {
        if product(n) % t == 0 {
            return n
        }
        n++
    }
}

func product(n int) int {
    var prod = 1
    for n > 0 {
        digit := n % 10
        prod*=digit
        n/=10
    }
    return prod
}
