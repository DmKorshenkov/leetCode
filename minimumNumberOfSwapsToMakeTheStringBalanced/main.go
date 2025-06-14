func minSwaps(s string) int {
    swap := 1
    closedOpen := 0
    for _, r := range s {
        if r == '[' {
            closedOpen++
        } else {
            closedOpen--
            if closedOpen == -1 {
                swap++
                closedOpen = 0
            }
        }
    }

    return swap / 2
}


func threeSum(nums []int) [][]int {
    resM := make(map[[3]int]byte)
    mp := make(map[int]int)
    for i := 1; i < len(nums); i++ {
        mp[nums[i]] = i
    }

    for i := 0; i < len(nums); i++ {
        for j := i+1; j < len(nums); j++ {
            K := nums[i] + nums[j]
            if k, ok := mp[-K]; ok && k > j { 
                arr := [3]int{nums[i],nums[j],nums[k]}
                slice := arr[:] 
                sort.Ints(slice)            
                resM[arr]++
            }
        }
    }

    res := [][]int{}
    for key := range resM {
        res = append(res, key[:])
    }
    return res
}
