func minOperations(logs []string) int {
    ans := []byte{}
    for i := range logs {
        if len(ans) > 0 && logs[i] == "../" {
            ans = ans[:len(ans)-1]
            continue
        }

        if logs[i] != "./" && logs[i] != "../" {
            ans = append(ans, 1)
        }
    }
    return len(ans)
}
