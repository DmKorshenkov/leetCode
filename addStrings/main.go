package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(addStrings("1999", "1000"))

}

func addStrings(num1 string, num2 string) string {

	bb := bytes.Buffer{}
	b2 := bytes.Buffer{}
	n := int32(0)

	if len(num1) > len(num2) {
		max := len(num1) - 1
		min := len(num2) - 1

		for _ = 0; min > -1; min-- {
			n += int32((num1[max] - 48) + (num2[min] - 48))
			bb.WriteRune((n % 10) + 48)
			n = n / 10
			max--
		}
		for max > -1 {
			fmt.Println(n)
			n += int32(num1[max] - 48)
			bb.WriteRune((n % 10) + 48)
			n = n / 10
			max--
		}
		if n == 1 {
			bb.WriteRune('1')
		}
	} else {
		max := len(num2) - 1
		min := len(num1) - 1

		for _ = 0; min > -1; min-- {
			n += int32((num2[max] - 48) + (num1[min] - 48))
			bb.WriteRune((n % 10) + 48)
			n = n / 10
			max--
		}
		for max > -1 {
			fmt.Println(n)
			n += int32(num2[max] - 48)
			bb.WriteRune((n % 10) + 48)
			n = n / 10
			max--
		}
		if n == 1 {
			bb.WriteRune('1')
		}
	}

	for in := bb.Len() - 1; in > -1; in-- {
		b2.WriteByte(bb.Bytes()[in])
	}

	return b2.String()
}
