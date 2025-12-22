func subsests(nums []int) [][]int {
	var res [][]int
	var path = make([]int, 0, len(nums))

	var back func(int)

	back = func(start int){
		tmp := make([]int, len([path))
		copy(tmp, path)
		res = append(res, tmp)

		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])

			back(i+1)

			path = path[:len(path)-1]
		}
	}
	back(0)
	return res
}
