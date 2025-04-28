func uniqueOccurrences(arr []int) bool {
	mp := make(map[int]int)

	for i := range arr {
		if _, ok := mp[arr[i]]; ok {
			mp[arr[i]]++
		} else {
			mp[arr[i]] = 1
		}
		//		fmt.Println(arr[i])
	}
	for key, val := range mp {
		for key2, val2 := range mp {
			if key != key2 && val == val2 {
				return false
			}
		}
	}
	return true
}
