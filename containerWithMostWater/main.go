func maxArea(height []int) int {
    var res int
    var tmp int
    
    res = 0
    tmp = 0
    for i, j := 0, len(height)-1; i <= j; {

        tmp = min(height[i], height[j]) * (j-i)

        if tmp > res {
            res = tmp
        }

        if height[i] < height[j] {
            i++
        } else {
            j--
        }

    }  

  return res
}
