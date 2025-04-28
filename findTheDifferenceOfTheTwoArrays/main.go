func findDifference(nums1 []int, nums2 []int) [][]int {
	mp := make(map[int]int)
	res := make([][]int, 2)
//	fmt.Println(nums1, nums2)
	for i := range nums1 {
		mp[nums1[i]]++
	}

	for i := range nums2 {
		if _, ok := mp[nums2[i]]; ok {
			mp[nums2[i]] = -1
		} else {
            mp[nums2[i]] = 0
			res[1] = append(res[1], nums2[i])

		}
	}

	for key, val := range mp {
		if val > 0 {
			res[0] = append(res[0], key)
		}
	}
//	fmt.Println(res[0], res[1])
	return res
}
