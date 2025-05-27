func minLength(s string) int {
    stack := []rune{}

    for _, i := range s {
        if len(stack) > 0 && (i == 'B' || i == 'D') {
            if i == 'B' && stack[len(stack)-1] == 'A' {
                stack = stack[:len(stack)-1]
                continue
            }
            if i == 'D' && stack[len(stack)-1] == 'C' {
                stack = stack[:len(stack)-1]
                continue
            }
        }
        
        stack = append(stack, i)
        
    }
    return len(stack)
}
