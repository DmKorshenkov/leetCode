func intersection(nums1 []int, nums2 []int) []int {
    arr := make([]int, 0,len(nums1)) 
    for in := range nums1{
        
        for in2 := range nums2 {
            if nums1[in] == nums2[in2]{
                ok := false
                for in3 := range arr{
                    if nums1[in] == arr[in3]{
                        ok = true
                    }
                }
                if !ok {
                arr = append(arr, nums1[in])
                }
            }
        }
    }
    return arr
}
