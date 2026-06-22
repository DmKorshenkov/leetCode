func maxIceCream(costs []int, coins int) int {
	max := costs[0]

	for i := 0; i < len(costs); i++ {
		if costs[i] > max {
			max = costs[i]
		}
	}

	count := make([]int, max+1)

	for _, n := range costs {
		count[n]++
	}
	
	res := 0
    for i := 1; i < len(count); i++ {
		if count[i] == 0 {
			continue
		}
		if count[i] * i < coins {
			res += count[i]
			coins -= count[i] * i
		} else {
			res += coins / i
            break
		}

	}
	return res
}
