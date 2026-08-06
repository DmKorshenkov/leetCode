func findingUsersActiveMinutes(logs [][]int, k int) []int {
	uam := make([]int, k)
	hash := make(map[int]map[int]struct{})
	for i := 0; i < len(logs); i++ {
		ID := logs[i][0]
		time := logs[i][1]
		if _, ok := hash[ID]; !ok {
			hash[ID] = make(map[int]struct{})
		}
		hash[ID][time] = struct{}{}
	}

	for ID := range hash {
		uam[len(hash[ID])-1]++
	}

	return uam
}
