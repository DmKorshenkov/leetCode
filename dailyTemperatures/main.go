func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	stack := make([]int, 0, len(temperatures))

	for i := len(temperatures) - 1; i >= 0; i-- {

		for len(stack) > 0 && temperatures[stack[len(stack)-1]] <= temperatures[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			ans[i] = stack[len(stack)-1] - i
		}

		stack = append(stack, i)
	}

	return ans
}
