package main

func main() {

}

func minWindow(s string, t string) string {
	mpT := make(map[byte]int)
	sS := make([]int, 128)
	var res = make([]byte, len(s)+1)
	for i := range t {
		mpT[t[i]]++
	}

	start := 0
	for end := 0; end < len(s); end++ {

		sS[s[end]]++
		for eq3(sS, mpT) && start <= end {
			if (end-start)+1 < len(res) {
				res = []byte(s[start : end+1])
			}
			sS[s[start]]--
			start++
		}
//		fmt.Println(s[start : end+1])

	}
	if len(res) == len(s)+1 {
		//		fmt.Println("1")
		return ""
	}
	return string(res)
}

func eq3(window []int, mp map[byte]int) bool {
	for c := range mp {
		if mp[c] > window[c] {
			return false
		}
	}
	return true
}
