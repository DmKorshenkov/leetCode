func minRemoveToMakeValid(s string) string {
    stack := []int{}

    for i, r := range s {
        if len(stack) > 0 && r == ')' {
            if s[stack[len(stack)-1]] == byte('(') {
                stack = stack[:len(stack)-1]
                continue
            }
        }
        if r == '(' || r == ')' {
            stack = append(stack, i)
        }
    }

    bb := bytes.Buffer{}

    for i := range s {
        if len(stack) > 0 && i == stack[0] {
            stack = stack[1:]
            continue
        }
        bb.WriteByte(s[i])
    }
    return bb.String()
}
