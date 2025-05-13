func findDuplicate(nums []int) int {
  arr := make([]int, len(nums))  

  for i := range nums {
    arr[nums[i]]++
    if arr[nums[i]] > 1 {
        return nums[i]
    }
  }
  return 0
}
