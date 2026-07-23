func xorOperation(n int, start int) int {
    var xor int
    for i := 0; i < n; i++ {
        xor ^= start + (2 *i)
    }
    return xor
}
