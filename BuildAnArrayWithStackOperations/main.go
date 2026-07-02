func buildArray(target []int, n int) []string {
    answer := []string{}
    step := 0
    for i := range target {
        for  {
            answer = append(answer, "Push")
            step++
            if step == target[i] {
                break
            }
            answer = append(answer, "Pop")
        }
    }
    return answer
}
