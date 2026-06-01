func minimumCost(cost []int) int {
    sort.Ints(cost)
    var sum int
    for i, step := len(cost)-1, 0; i >= 0; i-- {
        if step == 2 {
            step = 0
            continue
        } 
        sum += cost[i]
        step++
    }
    return sum
}
