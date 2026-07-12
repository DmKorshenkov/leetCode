func arrayRankTransform(arr []int) []int {
	ln := len(arr)
	tempo := append(make([]int, 0, ln), arr...)
	sort.Slice(tempo, func(i, j int) bool { return tempo[j] > tempo[i] })
	//	fmt.Println(tempo)
	hashUniq := make(map[int]int)
	for i, sep := 0, 1; i < ln; i++ {
		if _, ok := hashUniq[tempo[i]]; ok {
			sep--
		} else {
			hashUniq[tempo[i]] = i + sep
		}
	}
	//	fmt.Println(hashUniq)
	for i := 0; i < ln; i++ {
		arr[i] = hashUniq[arr[i]]
	}
	//	fmt.Println(ok(arr))
	return arr
}
