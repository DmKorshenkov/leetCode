func arrayStringsAreEqual(word1 []string, word2 []string) bool {
    bb := bytes.Buffer{}
    b2 := bytes.Buffer{}
    
    for i := 0; i < len(word1); i++ {
        bb.WriteString(word1[i])
    }
    word1 = nil
    for i := 0; i < len(word2); i++ {
        b2.WriteString(word2[i])
    }
    word2 = nil
    return bb.String() == b2.String()
}
