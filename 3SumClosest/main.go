func threeSumClosest(nums []int, target int) int {
    answer := nums[0]+nums[1]+nums[2]
    minDiff := abs(answer - target)
    for a := 0; a < len(nums)-2; a++ {
        for b := a+1; b < len(nums)-1; b++ {
            for c := b+1; c < len(nums); c++ {
                tmp := nums[a]+nums[b]+nums[c]
                diff := abs(tmp - target)
                if diff < minDiff {
			        minDiff = diff
			        answer = tmp
		        }
            }
        }
    }

	return answer
}

func abs(n int) int {
    if n < 0 {
        return -n
    }

    return n
}

