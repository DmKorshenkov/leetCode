func countConsistentStrings(allowed string, words []string) int {
    mp := make(map[byte]uint8)  
    for i := range allowed {
        mp[allowed[i]] = 1
    }
    var ok bool
    var res int

    res = 0
    ok = false
    for iWords := range words {
        ok = true
        for iWord := range words[iWords] {
            if _, check := mp[words[iWords][iWord]]; !check {
                ok = false
                break
            }
        }
        if ok {
            res++
        }
    }
    return res
}
