func reverseStr(s string, k int) string {
    bb := bytes.Buffer{}
    bb.WriteString(s)

    for i := 0; i < bb.Len(); i = i + (k*2) {
        if i + k < len(s) {
            reverse(bb.Bytes()[i : (i + k)])
        } else {
            reverse(bb.Bytes()[i : ])
        }
    }
    return bb.String()
}

func reverse(str []byte) {
	for i, j := 0, len(str)-1; i <= j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
}
