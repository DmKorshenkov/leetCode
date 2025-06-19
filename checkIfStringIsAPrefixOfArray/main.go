func isPrefixString(s string, words []string) bool {

    bb := bytes.Buffer{}
    for i := range words {
        bb.WriteString(words[i])
        if len(bb.String()) >= len(s) {
            break
        } 
    }
//    min := min(len(s), len(bb.Bytes()))
    return s == bb.String()
}
