func removeDuplicates(s string) string {
    stack := make([]rune, 0, len(s))
    for _, b := range s {
        if len(stack) > 0 && b == stack[len(stack)-1] {
            stack = stack[:len(stack)-1]
        } else {
            stack = append(stack, b)
        }
    }
    return string(stack)
}
