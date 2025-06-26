func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
    var res int

    res = 0
    mp1 := make(map[int]int)

    for a := 0; a < len(nums1); a++ {
        for b := 0; b < len(nums2); b++ {
            mp1[nums1[a]+nums2[b]]++
        }
    }

    for c := 0; c < len(nums3); c++ {
        for d := 0; d < len(nums4); d++ {
            if val, ok := mp1[-(nums3[c]+nums4[d])]; ok {
                res += val
            }
        }
    }
    return res
}
