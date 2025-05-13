func reverseWords(s string) string {
    
    bS := []byte(s)
    bb := bytes.Buffer{}
    buff := []byte{}
    for i := range bS {
        if bS[i] == byte(' '){
            bb.WriteString(string(buff) + string(' '))
            buff = nil
            continue
        }
        buff = append([]byte{bS[i]}, buff...)
    }
    bb.WriteString(string(buff))
    return bb.String()
}
