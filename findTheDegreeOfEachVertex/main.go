func findDegrees(matrix [][]int) []int {
    n := len(matrix)
 //   ans := make([]int, n)
    for j := 0; j < n; j++ {
        for i := 0; i < n-1; i++ {
            matrix[n-1][j]+=matrix[i][j]
        }
    }
    return matrix[n-1]
}
