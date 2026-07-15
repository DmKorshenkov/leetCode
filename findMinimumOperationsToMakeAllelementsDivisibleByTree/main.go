func minimumOperations(nums []int) int {
    var count int
    for _, n := range nums {
        if n % 3 != 0 {
            count++
        }
    }
    return count
}
