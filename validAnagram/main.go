package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(isAnagram("gifla", "figal"))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sLen := len(s)
	var sumS = (make([]int, len(s), len(s)))
	var sumT = (make([]int, len(s), len(s)))
	for in := 0; in < sLen; in++ {
		sumS[in] = int((s[in]))
		sumT[in] = int((t[in]))
	}
	sort.Ints(sumS)
	sort.Ints(sumT)
	for in := range sumS {
		if sumS[in] != sumT[in] {
			return false
		}
	}
	//	fmt.Println(sumS, sumT)
	return true
}
