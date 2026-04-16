func closestTarget(words []string, target string, startIndex int) int {
	if words[startIndex] == target {
        return 0
    }
    n := len(words)
	step := 0
	for i, j := startIndex, startIndex; step < len(words); i++ {
		if words[(i+1)%n] == target || words[(j-1+n)%n] == target {
			return step + 1
		}
		j--
		step++
	}
	return -1
}

func init() {
    debug.SetMemoryLimit(0)
}
