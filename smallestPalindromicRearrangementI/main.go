func smallestPalindrome(s string) string {
 
    hash := make([]int, 26)
    for i := 0; i < len(s); i++ {
        hash[s[i]-97]++
    }
    palindrom := make([]byte, len(s))
    
    start, end := 0, len(s)-1

    for i := 0; i < 26; i++ {
        for hash[i] >= 2 {
            c := byte(i+97)
            palindrom[start], palindrom[end] = c,c
            start++
            end--
            hash[i] -= 2
        } 
        if hash[i] == 1 {
            palindrom[len(s)/2] = byte(i+97)
        }
    }

    return string(palindrom)
}
