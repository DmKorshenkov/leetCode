func findLength(nums1 []int, nums2 []int) int {
    m := len(nums1)
    n := len(nums2)

    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }
    res := 0
    for i := 1; i <= len(nums1); i++ {
        for j := 1; j <= len(nums2); j++ {
            if nums1[i-1] == nums2[j-1] {
                dp[i][j] = dp[i-1][j-1] + 1
                if dp[i][j] > res {
                    res = dp[i][j]
                }
            } else {
                dp[i][j] = 0
            }
        }
    }
    return res
}
