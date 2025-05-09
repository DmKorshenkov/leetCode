func pivotArray(nums []int, pivot int) []int {
    l := make([]int, 0,len(nums))
    r := make([]int, 0,len(nums))
    for i := 0; i < len(nums); i++{
        if nums[i] < pivot {
            l = append(l, nums[i])
        }
        if nums[i] > pivot {
            r = append(r, nums[i])
        }
    }
    m := len(nums) - (len(l) + len(r))
    for i := 0; i < m; i++ {
        l = append(l, pivot)
    }
  return append(l, r...)   
}
