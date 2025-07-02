func searchMatrix(matrix [][]int, target int) bool {
    for j := len(matrix[0])-1; j >= 0; j--{
        if matrix[0][j] <= target {
            
            for i := 0; i < len(matrix); i++{
                if matrix[i][j] == target {
                    return true
                }
            }
        }
    }

    return false
}
