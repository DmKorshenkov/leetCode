func removeStars(s string) string {
    stack := make([]byte, 0, len(s))

    for i := 0; i < len(s); i++ {
        if s[i] != '*' {
            stack = append(stack, s[i])
        }

        if s[i] == '*' {
            stack = stack[:len(stack)-1]
        }
    }
//    fmt.Println(stack)
    return string(stack)
}
