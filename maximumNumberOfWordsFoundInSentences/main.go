func mostWordsFound(sentences []string) int {
    count := 1
    for i := len(sentences)-1; i >= 0; i-- {
        countSpace := 1
        for c := range sentences[i] {
            if sentences[i][c] == ' ' {
                countSpace++
            }
        }
        count = max(count, countSpace)
    }
    return count
}
