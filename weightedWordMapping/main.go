func mapWordWeights(words []string, weights []int) string {
    result := make([]byte, len(words))
    for i := 0; i < len(words); i++ {
        sum := 0
        for j := 0; j < len(words[i]);j++ {
            sum += weights[int(words[i][j] - 97)]
        }
        result[i] = byte(122 - (sum % 26))
    }
    return string(result)
}
