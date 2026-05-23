func check(nums []int) bool {
    lenght, ok := len(nums), false
    A := append(make([]int, 0,lenght), nums...)
    sort.Slice(nums, func(i, j int)bool{return nums[i] < nums[j]})
    for x := 0; x < len(nums); x++ {
        for i := range nums {
            if nums[i] == A[(i+x) % lenght] {
                ok = true
            } else {ok = false; break}

        }
        if ok {
            return true
        }
    }

    return false
}
