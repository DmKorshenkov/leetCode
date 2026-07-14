func digitFrequencyScore(n int) int {
    score := 0
    for n > 0 {
        score += n%10
        n/=10
    }
    return score
}
