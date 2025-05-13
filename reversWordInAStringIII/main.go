func reverseWords(s string) string {
    
    sliceS := strings.Split(s, " ")

    for i := range sliceS {
        sliceS[i] = reverse([]byte(sliceS[i]))
    }
    return strings.Join(sliceS, " ")
}

func reverse(str []byte) string{
	for i, j := 0, len(str)-1; i <= j; i, j = i+1, j-1 {
        str[i],str[j] = str[j],str[i]
    }
    return string(str)
}


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
