package main

import (
	"strings"
)

func main() {

}

func longestCommonPrefix(strs []string) string {
	var sb strings.Builder
	if len(strs) == 1 {
		return strs[0]
	}

	var ok = true
	for in := 0; in < len(strs[0]); in++ {

		for min := 1; min < len(strs) && ok; min++ {
			ok = func() bool {
				if in+1 > len(strs[min]) {
					return false
				}
				if strs[min][in] == strs[0][in] {
					return true
				}
				return false
			}()
		}
		if ok == false {
			return sb.String()
		}
		sb.WriteByte(strs[0][in])
	}
	return sb.String()
}
