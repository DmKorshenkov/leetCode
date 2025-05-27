func buildArray(target []int, n int) []string {
    ans := make([]string, 0, n)
    stack := []int{}
   
    for i, num := 0, 1; i < len(target); num++{
       
        stack = append(stack, num)
        ans = append(ans, "Push")

        if stack[len(stack)-1] != target[i] {
            stack = stack[:len(stack)-1]
            ans = append(ans, "Pop")
            continue
        } else {
            i++
        }

    }
    return ans
}
