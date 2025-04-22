func mergeAlternately(word1 string, word2 string) string {
    max := len(word1)

    if len(word2) > len(word1) {
        max = len(word2)
    }
    bb := bytes.Buffer{}

    for i := 0; i < max; i++{
        if i < len(word1) {
            bb.WriteByte(word1[i])
        }
        if i < len(word2) {
            bb.WriteByte(word2[i])
        }
 
    }
    return bb.String()
}
