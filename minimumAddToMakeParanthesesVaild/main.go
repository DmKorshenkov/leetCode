func minAddToMakeValid(s string) int {
    stack := []rune{}

    for _, r := range s {
        if len(stack) > 0 {
            if r == ')' && stack[len(stack)-1] == '(' {
                stack = stack[:len(stack)-1]
                continue
            }
        }

        stack = append(stack, r)
    }
//    fmt.Println(len(stack))
    return len(stack)
}
