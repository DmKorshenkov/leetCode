func reversePrefix(s string, k int) string {
    reverse := []byte(s)

    for l, r := 0, k-1; l < r; {
        reverse[l], reverse[r] = reverse[r],reverse[l]
        l++
        r--
    }
    return string(reverse)
}
