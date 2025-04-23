package main

import (
	"bytes"
	"fmt"
)

func main() {
	helper("hello world suck")
	//	str := "123 456"
	/*	arr := []byte{}
		for i := range str {
			arr = append(arr, str[i])
		}
		fmt.Println(str)
		reverseString(arr[3:])
		fmt.Println(string(arr))*/
	helper("the   sky   is  blue")

}

func helper(s string) {

	bb := bytes.Buffer{}
	bb.WriteString(s)
	str := make([]byte, len(s))
	for i := 0; i < bb.Len(); i++ {
		str[i] = bb.Bytes()[i]
	}
	ln := len(s) - 1
	//	reverseString(str)
	//	fmt.Println(string(str))
	for i := 0; i <= ln; i++ {
		//		fmt.Println(string(str))
		//		time.Sleep(time.Second * 3)
		//		if str[i] == byte(' ') {
		//		}
		if str[ln] == byte(' ') {
			reverseString(str[i : len(str)-ln-1])
			//			reverseString(str[:ln])
		}
		//		if str[i] == byte(' ') {
		//			reverseString(str[i+1:])
		//		}
		str[i], str[ln] = str[ln], str[i]
		ln--

		//		fmt.Println(i, len(str)-ln)
	}
	start := 0
	for i := 0; i < len(str); i++ {

		if str[i] == byte(' ') {
			reverseString(str[start:i])
			start = i + 1

		}
		if i == len(str)-1 {
			reverseString(str[start:])
		}
	}
	fmt.Println(string(str))
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
