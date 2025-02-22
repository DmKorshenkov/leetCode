func findErrorNums(nums []int) []int {
   sort.Ints(nums) 
   res := 0
   for in := range nums{   
    
    res+=in+1
    res-=nums[in]
   }
   for in := 0; in < len(nums)-1; in++{
        if nums[in] == nums[in+1]{
            return[]int{nums[in],res+nums[in]}
        }
   }
   return nil
}
