func numberGame(nums []int) []int {

  sort.Ints(nums)

  for i := 0; i < len(nums)-1; i+=2{
    nums[i],nums[i+1] = nums[i+1],nums[i]
  } 

  return nums
}
