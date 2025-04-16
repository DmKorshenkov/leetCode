func defangIPaddr(address string) string {
    bb := bytes.Buffer{}

    for i := 0; i < len(address); i++{
        if address[i] == 46 {
            bb.WriteString("[.]")
        } else {
            bb.WriteByte(address[i])
        }
    }
    return bb.String()
}
