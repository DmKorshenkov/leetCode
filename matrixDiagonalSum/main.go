func diagonalSum(mat [][]int) int {
	var res int

	n := len(mat)
	if n == 1 {
		return mat[0][0]
	}

	if n % 2 != 0 {
		res -= mat[n/2][n/2]
	}
    
	for end := 0; end < n; end++ {
		res += mat[end][end]
	}
    n--
	for end := 0; end < len(mat); end++ {
		res += mat[end][n]
		n--
	}


	return res
}
