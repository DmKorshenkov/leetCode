func customSortString(order string, s string) string {
    bb := bytes.Buffer{}
    res := bytes.Buffer{}

    bb.WriteString(s)
    for i := range order {
        for iB := 0; iB < bb.Len(); iB++ {
            if order[i] == bb.Bytes()[iB] {
                res.WriteByte(bb.Bytes()[iB])
                bb.Bytes()[iB] = '0'
            }
        }
    }

    for i := range bb.Bytes() {
        if bb.Bytes()[i] != byte('0') {
            res.WriteByte(bb.Bytes()[i])
        }
    }
    return res.String()
}
