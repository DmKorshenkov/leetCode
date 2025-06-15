func containsNearbyDuplicate(nums []int, k int) bool {
  for i := 0; i < len(nums); i++ {

    for  j := i+1; j < len(nums) && abs(i-j) <= k; j++ {
        if nums[i] == nums[j] {
            return true
        }
    }
  }  

  return false
}

func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}
