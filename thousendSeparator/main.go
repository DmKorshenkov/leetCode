func thousandSeparator(n int) string {
    res := bytes.Buffer{}

    str := strconv.Itoa(n)
    str = reverse([]byte(str))
    for i := len(str)-1; i >= 0; i--{
        res.WriteByte(str[i])
        if i % 3 == 0 && i != 0 {
            res.WriteRune('.')
        }
    }
    return res.String()
}

func reverse(str []byte) string{
    for i, j := 0, len(str)-1; i < j; i,j=i+1,j-1{
        str[i],str[j]=str[j],str[i]
    }
    return string(str)
}
