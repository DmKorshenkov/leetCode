package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(addBinary("1111", "1111"))
	//	fmt.Println(int('1'), int('0'))
	//	fmt.Println(0 * degree(2, 1))
	// fmt.Println(x)
}

func addBinary(a string, b string) string {

	bb := bytes.Buffer{}
	lenB := len(b) - 1
	lenA := len(a) - 1
	max := len(b)

	if len(a) > len(b) {
		max = len(a)
	}

	arr := make([]int, max+1)

	for i := 0; i < len(arr); i++ {

		if lenA >= 0 {
			arr[i] += int(a[lenA]) - 48
			lenA--
		}
		if lenB >= 0 {
			arr[i] += int(b[lenB]) - 48
			lenB--
		}

	}
	//	fmt.Println(arr)

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == 2 {
			arr[i] = 0
			arr[i+1]++
		}
		if arr[i] == 3 {
			arr[i] = 1
			arr[i+1]++
		}
	}

	//	fmt.Println(arr)

	for i := len(arr) - 1; i > -1; i-- {

		bb.WriteRune(rune(arr[i] + 48))
	}
	if bb.String()[0] == '0' {
		return bb.String()[1:]
	}
	return bb.String()
}
