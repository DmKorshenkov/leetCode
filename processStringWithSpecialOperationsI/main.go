func processStr(s string) string {
    result := make([]byte, 0, len(s))
    reverse := func(){
    
        for i, j := 0,len(result)-1; i < j; {
            result[i],result[j] = result[j], result[i]
            i++
            j--
        }
    
    }
    duplicates := func(){
        result = append(result, result...)
    }
    removes:= func(){
        if len(result) == 0 {
            return
        }
        result = result[:len(result)-1]
    }
    insert := func(c byte) {
        result = append(result, c)
    }

    for i := range s {
        switch s[i] {
            case '*':
                removes()
            case '#':
                duplicates()
            case '%':
                reverse()
            default:
                insert(s[i])
        }
       // fmt.Println(string(result))
    }
    return string(result)
}
