func countSubstrings(s string) int {
    res := 0

    for i := 0; i < len(s); i++ {
        res++
        for j := i+1; j < len(s); j++{
            if reverse(s[i:j+1]) {
                res++
            }
        }
    }
    return res
}

func reverse(s string) bool {
    sB := []byte(s)

    for i, j := 0, len(sB)-1; i < j; i, j = i+1, j-1 {
        sB[i],sB[j] = sB[j],sB[i]
    }

    return s == string(sB)
}
