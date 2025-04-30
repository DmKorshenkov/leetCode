func findPermutationDifference(s string, t string) int {
    var res int
    res = 0
    for iS := range s {
        for iT := range t {
            if s[iS] == t[iT] {
                res += max(iS,iT) - min(iS, iT)
                break
            }
        }
    }
	return int(res)
}

func findPermutationDifference(s string, t string) int {
    mp := make(map[byte]int)

    for iS := range s {
        mp[s[iS]] = iS
    }
    var res int

    res = 0
    for iT := range t {
        res += max(mp[t[iT]], iT) -  min(mp[t[iT]], iT)
    }
	return int(res)
}
