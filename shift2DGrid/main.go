func shiftGrid(grid [][]int, k int) [][]int {
	m := len(grid)
	n := len(grid[m-1])
	ln := m * n
	vector := make([]int, 0, ln)
	for i := 0; i < m; i++ {
		vector = append(vector, grid[i]...)
	}
	//end
	for i, end, step := 0, k%ln, 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if end > 0 {
				grid[i][j] = vector[ln-end]
				end--
			} else {
				grid[i][j] = vector[step]
				step++
			}
		}
	}
//	fmt.Println(grid)
	return grid
}
