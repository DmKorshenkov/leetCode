func backspaceCompare(s string, t string) bool {
    stackS := make([]byte, 0, 200)
    stackT := make([]byte, 0, 200)

    
    for i := 0 ; i < max(len(s), len(t)); i++ {
        if i < len(s) {
            if len(stackS) > 0 && s[i] == byte('#') {
                stackS = stackS[:len(stackS)-1]
            } else {
                if s[i] != byte('#') {
                    stackS = append(stackS, s[i])
                }
            }

        }

        if i < len(t) {
            if len(stackT) > 0 && t[i] == byte('#') {
                stackT = stackT[:len(stackT)-1]
            } else {
                if t[i] != byte('#') {
                    stackT = append(stackT, t[i])
                }
            }
        }
    }
 //   fmt.Println(string(stackS), string(stackT))
    return string(stackS) == string(stackT)
}
