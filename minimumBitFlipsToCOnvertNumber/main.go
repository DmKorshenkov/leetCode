func minBitFlips(start int, goal int) int {
    start ^= goal
    goal = 0
    for start > 0 {
        start = start & (start - 1)
        goal++
    }
    return goal

}
