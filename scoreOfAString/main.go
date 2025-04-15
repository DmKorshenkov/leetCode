func scoreOfString(s string) int {
    arr := make([]byte, len(s)-1)
 //   var max, min byte
    for i := 1; i < len(s); i++ {
       
        arr[i-1] =  s[i] - s[i-1] 
        if s[i-1] > s[i] {
            arr[i-1] = s[i-1] - s[i]
        }
    }
//    fmt.Println(arr)
    res := 0
    for i := range arr{
        res += int(arr[i])
    }
    return res
}
