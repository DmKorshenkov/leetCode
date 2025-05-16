func reverseStr(s string, k int) string {
    if len(s) <= k {
        return reverse([]byte(s))
    }
    bb := bytes.Buffer{}
    bb.WriteString(reverse([]byte(s[:k])))

        for i, j := k, 0; i < len(s); {
            if j != k {
                bb.WriteByte(s[i])
                j++
                i++
                continue
            } 
            if j == k {
                j = 0
                if i + k > len(s) {
                    bb.WriteString(reverse([]byte(s[i:])))
                    break
                }
                bb.WriteString(reverse([]byte(s[i:i+k])))
                i += k
                continue
            }
            i++
        }

    return bb.String()
}

func reverse(str []byte) string {

	for i, j := 0, len(str)-1; i <= j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
	return string(str)
}
