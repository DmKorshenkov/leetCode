func gcdSum(nums []int) int64 {
    n := len(nums)
    prefixGcd := make([]int, n)

    for i, mx := 0,0; i < n; i++ {
        if nums[i] > nums[mx] {
            mx = i
        }
        prefixGcd[i] = gcd(nums[i], nums[mx])
    }

    sort.Slice(prefixGcd, func(i,j int) bool {return prefixGcd[j] > prefixGcd[i]})
    res := int64(0)
    for i, j := 0,n-1; i < j; {
        res += int64(gcd(prefixGcd[i], prefixGcd[j]))
        i++
        j--
    }
    return res
}
func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a % b
    }
    
    return a
}
