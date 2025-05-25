func findMaxConsecutiveOnes(nums []int) int {
    var n, res int

    res = 0
    for i := 0; i < len(nums); i++ {
        n = 0
        
        if nums[i] == 1 {
            for j := i; j < len(nums) && nums[j] == 1; j++ {
                n++
            }
            res = max(res, n)
        }
   } 
   return res
}
