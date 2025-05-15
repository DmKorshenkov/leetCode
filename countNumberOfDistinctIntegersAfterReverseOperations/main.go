func countDistinctIntegers(nums []int) int {
   mp := make(map[int]struct{})

   for i := range nums {
    mp[nums[i]] = struct{}{}
    mp[reversNumber(nums[i])] = struct{}{}
   }

   return len(mp)
}

func reversNumber(n int) int {
	res := 0

	for n > 0 {
		digit := n % 10
		res = res*10 + digit
		n /= 10
	}

	return res
}
