package main

import (
	"bytes"
	"fmt"
	"sort"
)

func main() {
	//	fmt.Println(string('a' - 32))
	fmt.Println(sortVowels("lEetcOde"))
}

func sortVowels(s string) string {
	var check = "aeiou"
	sl := make([]byte, len(s))
	vowels := bytes.Buffer{}
	for r := range s {
		ok := true
		for i := range check {
			if s[r] == check[i] || s[r] == check[i]-32 {
				vowels.WriteByte(s[r])
				ok = false
				break
			}
		}
		if ok {
			sl[r] = s[r]
		}
	}
	sort.Slice(vowels.Bytes(), func(i, j int) bool {
		return vowels.Bytes()[i] < vowels.Bytes()[j]
	})

	step := 0
	for i := range sl {
		if sl[i] == 0 {
			sl[i] = vowels.Bytes()[step]
			step++
		}
		if step == vowels.Len() {
			break
		}
	}
	return string(sl)
}
