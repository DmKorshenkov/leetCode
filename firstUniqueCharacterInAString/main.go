func firstUniqChar(s string) int {
   ok := [26]int{}

   for _, r := range s {
        ok[r-97]++
   }

   for i, r := range s {
    if ok[r-97] == 1 {
        return i
    }
   }
   return -1
}


func firstUniqChar(s string) int {
    mp := make(map[byte]uint)
    for i := range s {
        ok := i
        for j := i+1; j < len(s); j++ {
            if s[i] == s[j] {
                ok = j
                mp[s[j]] = 1
                break
            }
        }
        if _, check := mp[s[i]]; ok == i && !check{
            return i
        }
    }
    return -1
}
