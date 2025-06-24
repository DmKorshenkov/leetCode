func findMinimumOperations(s1 string, s2 string, s3 string) int {
    maxPrefix := 0
    ln := min(len(s3),min(len(s1),len(s2)))
    for i := 0; i < ln; i++{
        if s1[i] == s2[i] && s3[i] == s1[i]{
            maxPrefix++
        } else {
            break
        }
    }
    if maxPrefix == 0 {
        return -1
    }
    
    return (len(s1)+len(s2)+len(s3))-(maxPrefix*3)
}
