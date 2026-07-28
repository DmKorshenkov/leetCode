unc numberOfMatches(n int) int {
    var match int
    for n > 1 {
        if n % 2 == 0 {
            match += n / 2
            n /= 2
        } else {
            match += n / 2
            n = (n/2)+1
        }
    }
    
    return match
}
