func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
    count := 0;
    for _, n := range hours {
        if n >= target {
         count++
        }
    }
    return count
}
