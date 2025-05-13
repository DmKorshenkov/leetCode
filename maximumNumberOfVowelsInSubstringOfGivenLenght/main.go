func maxVowels(s string, k int) int {
	vowel := "aeiou"
	//	bb := bytes.Buffer{}
	var res int

	res = 0
	for i := 0; i < k; i++ {
		if ok := strings.Contains(vowel, string(s[i])); ok {
			res++
		}

	}
	tmp := res
	for i, j := 1, k; j < len(s); i, j = i+1, j+1 {
		if ok := strings.Contains(vowel, string(s[i-1])); ok {
			tmp--
		}
		if ok := strings.Contains(vowel, string(s[j])); ok {
			tmp++
		}
		if tmp > res {
			res = tmp
		}
		if res == k {
			return res
		}
	}
//	fmt.Println(res)
	return res
}
