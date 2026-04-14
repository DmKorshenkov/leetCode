func rotate(matrix [][]int) {
	//m := len(matrix) - 1
	n := len(matrix)

	for j := 0; j < n/2; j++ {
		last := n - 1 - j

		for i := j; i < last; i++ {
			offset := i - j

			top := matrix[j][i]

			matrix[j][i] = matrix[last-offset][j]
			matrix[last-offset][j] = matrix[last][last-offset]
			matrix[last][last-offset] = matrix[i][last]
			matrix[i][last] = top
		}
	}

}
