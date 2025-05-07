func isSubsequence(s string, t string) bool {
    var ok bool
    var lastIn int
    
    ok = true
    lastIn = 0
    for is := 0; is < len(s); is++ {
        ok = false

        for it := lastIn; it < len(t); it++{

            if s[is] == t[it] {
                ok = true
                lastIn = it + 1
                break
            }
        }

        if !ok {
            return ok
        }
    }
    
    return ok
}
