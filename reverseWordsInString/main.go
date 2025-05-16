package main

import (
	"bytes"
	"strings"
)

func main() {

}

func reverseWords(s string) string {
	bb := bytes.Buffer{}
	arr := []string{}
	for i := len(s) - 1; i > -1; i-- {
		bb.WriteByte(s[i])
		if s[i] == byte(' ') || i == 0 {
			reverseString(bb.Bytes())
			if bb.String() != " " {
				arr = append(arr, strings.TrimSpace(bb.String()))
			}
			bb.Reset()
		}
	}
	//	fmt.Println(arr, len(arr))
	bb.Reset()
	for i := range arr {
		bb.WriteString(arr[i] + " ")
	}
	return string(bb.Bytes()[:bb.Len()-1])
}

func reverseString(s []byte) {
	sLen := len(s) - 1

	for in := 0; in <= sLen; in++ {
		buff := s[in]
		s[in] = s[sLen]
		s[sLen] = buff
		sLen--
	}
}
