func numberOfPairs(nums1 []int, nums2 []int, k int) int {
   var res int

   res = 0
   for i := 0; i < len(nums1); i++ {

        for j := 0; j < len(nums2); j++ {
            if (nums1[i] % (nums2[j] * k)) == 0 {
                res++
            }
        }
   } 
   return res
}
