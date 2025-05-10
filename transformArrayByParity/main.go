func transformArray(nums []int) []int {
   l := 0
   for i := 0; i < len(nums); i++ {
        if nums[i] % 2 == 0 {
            nums[l] = 0
            l++
        } 
    }

   for i := l; i < len(nums); i++ { nums[i] = 1}

   return nums
}
