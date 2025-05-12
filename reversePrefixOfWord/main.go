func reversePrefix(word string, ch byte) string {
    bb := bytes.Buffer{}
    var i int

    i = 0
    for ; i < len(word); i++{
        bb.WriteByte(word[i])
        if word[i] == ch {
                res := reversString(bb.String())
                
                if len(res) != len(word) {
                    res += word[i+1:]
                }
            return res
        }
    }

    return  word
}

func reversString(str string) string {
	if len(str) <= 1 {
		return str
	}
	return reversString(str[1:]) + string(str[0])
}
