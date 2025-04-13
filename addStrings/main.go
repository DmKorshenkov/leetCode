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
	max := len(num1)

	if len(num2) > len(num1) {
		max = len(num2)
	}
	var num byte
	bytes := make([]byte, max+1)
	for diff := 0; diff < max; diff++ {
		if len(num1)-1-diff > -1 {
			num += num1[len(num1)-1-diff] - 48
		}
		if len(num2)-1-diff > -1 {
			num += num2[len(num2)-1-diff] - 48
		}

		bytes[len(bytes)-1-diff] = num % 10

		if num/10 == 1 {
			bytes[len(bytes)-1-diff-1] = 1
			num = num / 10
		} else {
			num = 0
		}
	}

	//	fmt.Println(bytes)

	for i := 0; i < len(bytes); i++ {
		if i == 0 && bytes[i] == 0 {
			continue
		}

		bb.WriteByte(bytes[i] + 48)
	}

	return bb.String()
}
