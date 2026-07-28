func findCenter(edges [][]int) int {
    hash := make(map[int]int)
    for i := range edges {for j := range edges[i]{hash[edges[i][j]]++}}
    max := 0
    for i := range hash{
        if hash[i] > hash[max] {
            max = i
        }
    }
    return max
}
