func minMirrorPairDistance(nums []int) int {
    hash := make(map[int]int)

    var res = len(nums)
    for j := 0; j < len(nums); j++ {
        if i, ok := hash[nums[j]]; ok {
            res = min(res, j-i)
        }
        hash[reverse(nums[j])] = j
    }

    if res == len(nums) {
        return -1
    }
    return res
}

func reverse(n int)int {
    var rever int

    for n > 0 {
        rever *= 10
        rever += n % 10
        n/=10
    }
    return rever
}
