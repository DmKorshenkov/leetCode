func countMajoritySubarrays(nums []int, target int) int {
    count := 0
    ln := len(nums)
    for i := 0; i < ln; i++ {
        freq := 0        
        for j := i; j < ln; j++ {
            if nums[j] == target {
                freq+=1
                
            }
            if 2 * freq > j-i+1 {
                count++
            }
        }
    }

    return count
}
