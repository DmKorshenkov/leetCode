func numJewelsInStones(jewels string, stones string) int {
    hash := make(map[byte]struct{})
    for c := range jewels {
        hash[jewels[c]] = struct{}{}
    }
    var count int
    for i := 0; i < len(stones); i++ {
        if _, ok := hash[stones[i]]; ok {
            count++
        }
    }
    return count
}
