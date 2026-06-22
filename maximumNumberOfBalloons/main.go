func maxNumberOfBalloons(text string) int {

	hash := make(map[byte]int)

	for i := range text {
        hash[text[i]]++
	}
    
	return min(hash[98], hash[97], hash[108]/2, hash[111]/2, hash[110])
}
