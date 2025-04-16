func finalValueAfterOperations(operations []string) int {
    var sum int

    for i := range operations {
        if operations[i][1] == '+' {
            sum++
        } else {
            sum--
        }
    }
    return sum
}
