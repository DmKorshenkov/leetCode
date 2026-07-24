func largestCombination(candidates []int) int {
	//	maxBit := 0
	var hash [24]int 
	//count := 0
	for j := range candidates {
		for i := 0; i < bits.Len(uint(candidates[j])); i++ {
			bit := (candidates[j] >> i) & 1
			if bit == 1 {
				hash[i]++
			}
		}
	}
    var count int
    for i := 0; i < 24; i++ {
        count = max(count, hash[i])
    }
	return count
}
