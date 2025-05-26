func countWords(words1 []string, words2 []string) int {
    var ans int
    mp1 := make(map[string]int)
    mp2 := make(map[string]int)

    for i := range words1 {
        mp1[words1[i]]++
    }

    for i := range words2 {
        mp2[words2[i]]++
    }

    for key, _ := range mp1 {
        if mp1[key] == 1 && mp2[key] == 1 {
            ans++
        }
    }
    return ans
}

