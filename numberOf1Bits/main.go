func hammingWeight(n int) (count int) {
    str := fmt.Sprintf("%b",n)
    for i := range str {
        if str[i] == byte('1') {
            count++
        }
    }
    return count
}
