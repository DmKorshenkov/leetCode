func maxFreqSum(s string) int {
    vowel := map[byte]int{
        'a': 0,
        'e': 0,
        'i': 0,
        'o': 0,
        'u': 0,
    }
   
    consonant := make(map[byte]int)

    max1, max2 := 0,0
    for i := range s {
        if  _, ok := vowel[s[i]]; ok {
            vowel[s[i]]++
            max1 = max(max1, vowel[s[i]])
        } else {
            consonant[s[i]]++
            max2 = max(max2, consonant[s[i]])
        }
        
    }

   
    return max1+max2

}
