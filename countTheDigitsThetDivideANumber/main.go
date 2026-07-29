func countDigits(num int) int {
    n := num
    var count int
    for num > 0 {
        digit := num % 10
        if n % digit == 0 {
            count++
        }
        num /=10
    }
    return count
}
