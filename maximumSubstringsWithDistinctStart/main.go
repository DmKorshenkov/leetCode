func maxDistinct(s string) int {
    var hash [26]bool
    count := 0
    for i := range s {
        if !hash[s[i]-97] {
            hash[s[i]-97] = true
            count++
        }
    }
    return count
}
