func compress(chars []byte) int {
    bb := bytes.Buffer{}
    var amount int

    for i := 0; i < len(chars); i++ {
        amount = 0
        for j := i; j < len(chars) && chars[j] == chars[i]; j++{
            amount++
        }

        bb.WriteByte(chars[i])
        if amount > 1 {
            bb.WriteString(strconv.Itoa(amount))
        }
        i += amount-1
    }
    for i := range bb.Bytes() {
        chars[i] = bb.Bytes()[i]
    }
    return bb.Len()
}
