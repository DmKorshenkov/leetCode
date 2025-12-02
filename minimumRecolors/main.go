func minimumRecolors(blocks string, k int) int {
	wB := 0
	start := 0
	minWB := 0
	for in := 0; in < k; in++ {
		if blocks[in] == 'W' {
			wB++
		}
	}
	minWB = wB
	for end := k; end < len(blocks); end++ {
		if blocks[end] == 'W' {
			wB++
		}
		if blocks[start] == 'W' {
			wB--
		}
		fmt.Println(wB)
		minWB = min(minWB, wB)
		start++
	}
	return minWB
}
