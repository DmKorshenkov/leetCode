func firstMissingPositive(nums []int) int {
    sort.Slice(nums, func(i, j int) bool{return nums[i]>nums[j]})
    if nums[0] <= 0 {
        return 1
    }
    
    minMissis := 1
    for i := len(nums)-1; i >= 0;  {
        if nums[i] <= 0 {
            i--
            continue
        }
        // 3 3 1 1 1 0 -1
        if nums[i] > minMissis {
            return minMissis
        } else if nums[i] == minMissis{
            minMissis++ 
            i--
            continue
        } else if nums[i] < minMissis {
            i--
        }
    }

    return nums[0]+1
}
