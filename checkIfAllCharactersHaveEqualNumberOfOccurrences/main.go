func areOccurrencesEqual(s string) bool {
  mp := make(map[byte]int)

  for i := range s {
    mp[s[i]]++
  }  
    check := mp[s[0]]
  for _,val := range mp {
    if val != check {
        return false
    }
  }
    return true
}
