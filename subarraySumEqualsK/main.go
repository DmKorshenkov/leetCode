func subarraySum(nums []int, k int) int {
/*    prevSum := make([]int, len(nums)+1)

	var countSum int
	for i := 1; i <= len(nums); i++ {
		prevSum[i] = prevSum[i-1] + nums[i-1]
		for j := 0; j < i; j++ {

			if prevSum[i]-prevSum[j] == k {
                    countSum++
            }
//				fmt.Println(prevSum[i] - prevSum[j])
		}
	}
//	fmt.Println(prevSum)
	return countSum
    */
    hashSum := make(map[int]int)
    hashSum[0] = 1 
	countSum, sum := 0,0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
        
		if count, ok := hashSum[sum - k]; ok {
			countSum += count
		}
		hashSum[sum]++
	}
	return countSum
}
