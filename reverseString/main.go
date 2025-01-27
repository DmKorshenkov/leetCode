package main

func main() {

}
func reverseString(s []byte)  {
    sLen := len(s)-1

    for in := 0; in < sLen && in != sLen; in++{
        buff := s[in] 
        s[in] = s[sLen]
        s[sLen] = buff
        sLen--
    }
}
