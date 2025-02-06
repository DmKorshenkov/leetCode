package main

import (
	"bytes"
	"fmt"
)

func main() {
	help("1910", "91")
}

func help(num1 string, num2 string) {

	bb := bytes.Buffer{}
	//	b2 := bytes.Buffer{}
	n := int32(0)

	if len(num1) >= len(num2) {
		max := len(num1) - 1
		min := len(num2) - 1
		for _ = 0; min > -1; min-- {
			n += int32((num1[max] - 48) + (num2[min] - 48))
			bb.WriteRune((n % 10) + 48)
			n = int32((num1[max]-48)+(num2[min]-48)) / 10
			max--
		}
		fmt.Println(bb.String(), n)
		for max > -1 {
			n += int32(num1[max] - 48)
			bb.WriteRune((n % 10) + 48)
			n = (n - int32(num1[max]-48)) / 10
			max--
		}
		fmt.Println(bb.String(), n)
		if n == 1 {
			bb.WriteRune('1')
		}
		fmt.Println(bb.String())
	}

	// fmt.Println(b2.String())
}

func revers(sl []byte) string {
	bb := bytes.Buffer{}

	for in := len(sl) - 1; in > -1; in-- {
		bb.WriteByte(sl[in])
	}
	return bb.String()
}
