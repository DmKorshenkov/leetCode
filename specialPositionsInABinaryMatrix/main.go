package main

func main() {

}

func numSpecial(mat [][]int) int {
    var index int

    res := 0
    for m := 0; m < len(mat); m++ {
        countN := 0
        for n := 0; n < len(mat[m]); n++ {
            if mat[m][n] == 1 {
                index = n
                countN++
            }
        }
        if countN == 1 {
            countM := 0
            for i := 0; i < len(mat); i++ {
                if mat[i][index] == 1 {
                    countM++
                }
            }
            if countM == 1 {
                res++
            }
        }
    }
    return res
}
