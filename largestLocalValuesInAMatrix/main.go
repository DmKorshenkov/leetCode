package main

func main() {

}

func largestLocal(grid [][]int) [][]int {
	var n = len(grid)
	var ans = make([][]int, (n - 2))

	for i := 0; i < n-2; i++ {
		for j := 0; j < n-2; j++ {
			maximum := 0
			for i2 := i; i2 < i+3; i2++ {
				for j2 := j; j2 < j+3; j2++ {
					maximum = max(maximum, grid[i2][j2])
				}

			}
			ans[i] = append(ans[i], maximum)
		}
	}

	return ans
}
