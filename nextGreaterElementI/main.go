func nextGreaterElement(nums1 []int, nums2 []int) []int {
    ans := make([]int, len(nums1))
//    var ok bool
    for i := range nums1 {
        ok := false
        for j := range nums2 {
            if nums1[i] == nums2[j]{
                ans[i] = nums1[i]
                ok = true
            }
            if ok && nums2[j] > ans[i] {
                ans[i] = nums2[j]
                break
            }
        }
        if ans[i] == nums1[i] {
            ans[i] = -1
        }
    }

    return ans
}
