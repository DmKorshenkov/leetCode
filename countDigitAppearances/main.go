func countDigitOccurrences(nums []int, digit int) int {
    var count int
    for i := range nums {
        for nums[i] > 0 {
            if nums[i] % 10 == digit {
                count++
            }
            nums[i]/=10 
        }
    }
    return count
}
