func numJewelsInStones(jewels string, stones string) int {
    var sum int

    for i := range jewels{
        for j := range stones {
            if jewels[i] == stones[j] {
                sum++
            }
        }
    }
    return sum
}
