func leftRightDifference(nums []int) []int {
	prefix := make([]int, len(nums));

	for i := 1; i < len(nums); i++ {
		prefix[i] = prefix[i-1] + nums[i-1];
	}
	
    sum := 0;
	for i := len(nums) - 1; i >= 0; i-- {
		prefix[i] = abc(prefix[i] - sum);
        sum+=nums[i];
	}
	return prefix;
}

func abc(n int) int{
    if n <0{
        return -n
    }
    return n
}
