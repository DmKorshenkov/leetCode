func maxDepth(s string) int {
    stack := []byte{}
    ans := 0

    for i := range s {
        if s[i] == byte('(') {
            stack = append(stack, s[i])
        } 
        if s[i] == byte(')') {
            stack = stack[:len(stack)-1]
        }
        ans = max(ans, len(stack))
    }
    return ans
}
