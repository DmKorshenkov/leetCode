func thirdMax(nums []int) int {
   sort.Ints(nums)
   count := 0

   for in := len(nums)-1; in != 0; in--{
    if nums[in] > nums[in - 1]{
        count++
    } 
    if count == 2{
        return nums[in-1]
    }
   } 

   return nums[len(nums)-1]
}
