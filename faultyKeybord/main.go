func finalString(s string) string {
//    runtime.GC()
    bb := bytes.Buffer{}

    for i := range s {
        if s[i] != 'i' {
            bb.WriteByte(s[i])    
        } else {
            reverseS(bb.Bytes())
        }
    }
    return bb.String()
}

func reverseS(s []byte) {
    for i,j := 0, len(s)-1; i <= j; i,j = i+1, j-1 {
        s[i],s[j] = s[j],s[i]
    }
}
