func longestCommonPrefix(arr1 []int, arr2 []int) int {
	hash := make(map[int]struct{})
	for i := 0; i < len(arr1); i++ {
		for ;arr1[i] > 0; arr1[i] /= 10 {
            hash[arr1[i]] = struct{}{}
        }
	}
	res := 0
	for i := 0; i < len(arr2); i++ {
		for x := arr2[i]; x > 0; x /= 10 {
			if _, ok := hash[x]; ok {
                l := 0
                for ;x > 0; x /= 10 {
                    l++
                }
                res = max(res, l)
                break
            }
		}
	}
	return res
}
