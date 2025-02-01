package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(isPalindrome("a9"))
	// var b byte
}

func isPalindrome(s string) bool {
	bb := bytes.Buffer{}
	for in := range s {
		if s[in] <= 90 && s[in] >= 65 {
			bb.WriteByte(s[in] + uint8(32))
			continue
		}
		if s[in] <= 122 && s[in] >= 97 ||
			s[in] >= 48 && s[in] <= 57 {
			bb.WriteByte(s[in])
		}
	}
	s = bb.String()
	bb.Reset()
	for in := len(s) - 1; in > -1; in-- {
		bb.WriteByte(s[in])
	}

	return s == bb.String()
}
