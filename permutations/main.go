func permute(nums []int) [][]int {
	var result [][]int

	var backtrack func(path []int, used []bool)
	backtrack = func(path []int, used []bool) {
		if len(path) == len(nums) {
			// сохраняем копию текущего пути
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])

			backtrack(path, used)

			// откат
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	used := make([]bool, len(nums))
	backtrack([]int{}, used)

	return result
}
