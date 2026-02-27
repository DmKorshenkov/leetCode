package main

func main() {

}

func checkInclusion(s1 string, s2 string) bool {
	ln := len(s1)
	pS1 := make([]int, 26)
	pS2 := make([]int, 26)
	for i := range s1 {
		pS1[int(s1[i]-'a')]++
	}

	for i := 0; i < len(s2); i++ {
		c := int(s2[i] - 'a')
		pS2[c]++

		if i >= ln {
			pS2[int(s2[i-ln]-'a')]--
		}

		if pS2[c] > 0 && eq2(pS1, pS2) {
			return true
		}
	}
	return false
}

func eq2(arr1, arr2 []int) bool {
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
