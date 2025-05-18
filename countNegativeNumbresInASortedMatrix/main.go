func countNegatives(grid [][]int) int {
    
    var res int

    res = 0


    for i := len(grid)-1; i > -1; i-- {
//        fmt.Println(grid[1])
        fmt.Println(grid[i])
        for j := range grid[i] {
//            fmt.Println("!")
            if grid[i][j] >= 0 {
                continue
            } 
            res++
        }
    }

    return res
}
