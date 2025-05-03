package main

func main() {

}

func singleNumber(nums []int) int {
    var res int
    
    res = 0

    for _, n := range nums {
        res ^= n
    }
    return res
}



func singleNumber(nums []int) int {
   if len(nums) == 1 {
    return nums[0]
   }
    for now := range nums {
        ok := false
        for next := 0; next < len(nums) && !ok; next++{
            if nums[now] == nums[next] && next != now{
                ok = true
            }
        }
    if !ok {
        return nums[now]
    }
    }
    return nums[len(nums)-1]
}
