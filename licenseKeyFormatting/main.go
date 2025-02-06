package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(licenseKeyFormatting("4", 1))
	// fmt.Println(int32('a'), int32('z'))
	//	fmt.Println(0%4, 2%4, 3%4, 4%4, 5%4, 8%4)
}

func licenseKeyFormatting(s string, k int) string {

	bb := bytes.Buffer{}
	b2 := bytes.Buffer{}
	for in := len(s) - 1; in > -1; in-- {
		if s[in] == '-' {
			continue
		}
		if s[in] <= 122 && s[in] >= 97 {
			bb.WriteByte(s[in] - 32)
		} else {
			bb.WriteByte(s[in])
		}

	}
	for in := bb.Len() - 1; in > -1; in-- {
		b2.WriteByte(bb.Bytes()[in])
		if in%k == 0 && in != 0 {
			b2.WriteByte('-')
		}
	}
	return b2.String()
}
