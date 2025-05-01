func balancedStringSplit(s string) int {
    var res, check int

    res = 0
    check = 0
    for i := range s {
        if s[i] == byte('R') {
            check--
        }

        if s[i] == byte('L') {
            check++
        }

        if check == 0 {
            res++
        }
    }
   return res
}
