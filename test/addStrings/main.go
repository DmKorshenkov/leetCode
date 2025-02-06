package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(addString("10", "1"))

}

func addString(num1 string, num2 string) string {

	bb := bytes.Buffer{}
	b2 := bytes.Buffer{}
	n := int32(0)
	ln1 := len(num1)
	ln2 := len(num2)
	if ln1 >= ln2 {
		for in, in2 := ln2-1, ln2-1; in > -1; in-- {
			n += int32((num1[in] - 48) + (num2[in2] - 48))
			bb.WriteRune((n % 10) + 48)
			n = int32((num1[in]-48)+(num2[in]-48)) / 10
			in2--
		}
		fmt.Println(bb.String(), n)
		for in := ln2; in < len(num1); in++ {
			n += int32(num1[in] - 48)
			bb.WriteRune((n % 10) + 48)
			n = int32(num1[in]-48) / 10
		}
		//	fmt.Println(n)
		if n == 1 {
			bb.WriteRune('1')
		}
		//	fmt.Println(bb.String())
	}

	for in := bb.Len() - 1; in > -1; in-- {
		b2.WriteByte(bb.Bytes()[in])
	}

	return b2.String()
}
