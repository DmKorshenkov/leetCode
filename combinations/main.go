func combine(n int, k int) [][]int {
	var res [][]int
	var path = make([]int, 0, k)

	var back func(int)

	back = func(start int){
		if len(path) == k {
			tmp := make([]int, len(path)
			copy(tmp, path)
			res = append(res, tmp)
			tmp = nil
		}

		for start <= n {
			path = append(path, start)
			
			back(start+1)
			
			path = path[:len(path)-1]
			start++
		}
	}
	return res
}
