package main

import (
	"fmt"
)

func main() {
	fmt.Println(strStr("sad", "saddddd"))

}

func strStr(haystack string, needle string) int {
	for in := range haystack {
		if haystack[in] == needle[0] &&
			in+len(needle) <= len(haystack) &&
			haystack[in:in+len(needle)] == needle {
			return in

		}
	}
	return -1
}
