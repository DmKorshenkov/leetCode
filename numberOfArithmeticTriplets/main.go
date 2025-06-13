func arithmeticTriplets(nums []int, diff int) int {
    res := 0
    for i := 0; i < len(nums)-2; i++ {
        for j := i+1; j < len(nums)-1; j++ {
            if nums[i] + diff == nums[j] {
                for k := j+1; k < len(nums); k++{
                    if nums[j] + diff == nums[k] {
                        res++
                        break
                    }
                }
                break
            }
        }
    }
    return res
}


func arithmeticTriplets(nums []int, diff int) int {
    mp := make(map[int]int)
 //   mp[40000] = 0
    res := 0
    for i := 1; i < len(nums); i++ {
        mp[nums[i]] = i
    }
    
    for i := 0; i < len(nums); i++ {
        if j, ok := mp[nums[i]+diff]; ok && i < j {
            if k, ok := mp[nums[j]+diff]; ok && j < k {
              res++
              //  mp[40000]++
            }
        }
    }
    return res
}
