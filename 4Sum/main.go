func fourSum(nums []int, target int) [][]int {
    mp := make(map[[4]int]int)
    for a := 0; a < len(nums)-3; a++ {
        for b := a+1; b < len(nums)-2; b++ {
             for c := b+1; c < len(nums)-1; c++ {
                for d := c+1; d < len(nums); d++ {
                    if nums[a]+nums[b]+nums[c]+nums[d] == target {
                        
                        tmp := []int{nums[a],nums[b],nums[c],nums[d]}
                        sort.Ints(tmp)
                        //var arr [4]int
                        arr := [4]int(tmp[:])
                        if _, ok := mp[arr]; !ok {
                            mp[arr]++
                        }
                        break
                    }
                }       
            }
        }
    }

    res := [][]int{}
    for key, _ := range mp {
        res = append(res, []int{key[0], key[1], key[2], key[3]})
    }
    return res
}
