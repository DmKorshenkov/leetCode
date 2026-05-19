func shipWithinDays(weights []int, days int) int {
	help := func(n int) int {
		sum := 0
		count := 1
		for i := 0; i < len(weights); i++ {
			sum += weights[i]
			if sum > n {
				sum = weights[i]
				count++
			}
		}
		return count
	}
	minimal := weights[0]
	maximal := 0
	for i := 0; i < len(weights); i++ {
		minimal = max(weights[i], minimal)
		maximal += weights[i]
	}

	for minimal <= maximal {
		mid := minimal + (maximal-minimal)/2

		if help(mid) > days {
			minimal = mid + 1
		} else {
			maximal = mid - 1
		}
	}

	return minimal
}
