package main

import (
	"bytes"
	"fmt"
)

func main() {
	countAndSay(30)
}

func countAndSay(n int) string {
	var rle bytes.Buffer
	var answer bytes.Buffer

	rle.WriteRune('1')

	for i := 1; i < n; i++ {
		answer.Reset()
		answer.WriteString(duplet(rle.Bytes()))
		rle.Reset()
		rle.Write(answer.Bytes())
	}
	return rle.String()
}
func duplet(str []byte) string {
	var answer bytes.Buffer
	var count int
	var c byte
	count = 1
	c = str[0]
	for i := 1; i < len(str); i++ {
		if str[i] == str[i-1] {
			count++
		} else {
			answer.WriteString(n(count))
			answer.WriteByte(c)
			c = str[i]
			count = 1
		}
	}
	answer.WriteString(n(count))
	answer.WriteByte(c)
	return answer.String()
}

func n(str int) string {
	return fmt.Sprintf("%d", str)
}
