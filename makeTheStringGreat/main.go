func makeGood(s string) string {
    stack := []rune{}
    for _, b := range s {
        fmt.Println(b, byte(b))
        if (len(stack) > 0 && (b == stack[len(stack)-1]+32 || b +32 == stack[len(stack)-1])){
            stack = stack[:len(stack)-1]
        } else {
            stack = append(stack, b)
        }
    }
    return string(stack)
}
