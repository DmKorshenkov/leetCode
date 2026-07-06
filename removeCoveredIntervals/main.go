func removeCoveredIntervals(intervals [][]int) int {
    ln, count := len(intervals), 0
    for i := 0; i < ln; i++ {
        c, d := intervals[i][0], intervals[i][1]
        for j := 0; j < ln && c != -1 && d != -1; j++ {
            a, b := intervals[j][0], intervals[j][1]
            if a == -1 && b == -1 || i == j {
                continue
            }
            if c <= a && b <= d {
               // fmt.Println(i, j)
                count++
                intervals[j] = []int{-1,-1}
            }
        }
    }
    return ln - count
}
