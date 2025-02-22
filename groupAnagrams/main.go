package main

import "fmt"

func main() {
	fmt.Println(groupAnagrams([]string{"", ""}))

}
func groupAnagrams(strs []string) [][]string {
	//	runtime.GC()
	matrix := make([][]string, 0, len(strs))
	for now := 0; now < len(strs)-1; now++ {
		tmp := make([]string, 0, len(strs)/5)
		tmp = append(tmp, strs[now])
		for next := now + 1; next < len(strs) && strs[now] != "!"; next++ {
			if isAnagram2(strs[now], strs[next]) {
				tmp = append(tmp, strs[next])
				strs[next] = "!"
			}
		}
		if strs[now] != "!" {
			matrix = append(matrix, tmp)
		}
	}

	if strs[len(strs)-1] != "!" {
		matrix = append(matrix, []string{strs[len(strs)-1]})
	}
	return matrix
}

func isAnagram2(s string, t string) bool {
	if len(s) != len(t) || s == "!" || t == "!" {
		return false
	}
	array := [26]int{}
	for in := range s {
		array[s[in]-'a']++
		array[t[in]-'a']--
	}
	for in := range array {
		if array[in] != 0 {
			return false
		}
	}
	return true
}
