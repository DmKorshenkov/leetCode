func lengthOfLongestSubstring(s string) int {
	if len(s) == 1 {
		return 1
	}
	var res int

	res = 0
	mp := make(map[byte]struct{})

	for i := 0; i < len(s); i++ {
        if len(s) - i <= res {
            return res
        }
		//		mp[s[i]] = struct{}{}
		for j := i; j < len(s); j++ {
			if _, ok := mp[s[j]]; ok {
				break
			} else {
				mp[s[j]] = struct{}{}

			}

		}

		if len(mp) > res {
			res = len(mp)
		}
		mp = nil
		mp = make(map[byte]struct{})
	}

	if res < len(mp) {
		res = len(mp)
	}
//	fmt.Println(res)
	return res
}
