func transpose(matrix [][]int) [][]int {
    var m = len(matrix)
	var n = len(matrix[m-1])
	var res = make([][]int, n)

	for i := 0; i < n; i++ {
        res[i] = make([]int, m)
		for j := 0; j < m; j++ {
			res[i][j] = matrix[j][i]
		}
	}
	return res
}
