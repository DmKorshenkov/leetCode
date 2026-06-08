type NumMatrix struct {
    pref [][]int
}


func Constructor(matrix [][]int) NumMatrix {
    	col := len(matrix)
	row := len(matrix[0])

	pref := make([][]int, col+1)
	pref[0] = make([]int, row+1)
	for i := 1; i <= col; i++ {
		pref[i] = make([]int, row+1)
		for j := 1; j <= row; j++ {
			pref[i][j] =
				matrix[i-1][j-1] +
					pref[i-1][j] +
					pref[i][j-1] -
					pref[i-1][j-1]
		}
	}
    return NumMatrix{pref: pref}
}


func (this *NumMatrix) SumRegion(r1 int, c1 int, r2 int, c2 int) int {
    return this.pref[r2+1][c2+1]-this.pref[r1][c2+1]-this.pref[r2+1][c1]+this.pref[r1][c1]
}
