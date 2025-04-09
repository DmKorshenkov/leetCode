package main

import (
	"fmt"
)

func main() {
	fmt.Println(addBinary("1010", "111"))
	//	fmt.Println(int('1'), int('0'))
	//	fmt.Println(0 * degree(2, 1))
	// fmt.Println(x)
}

func addBinary(a string, b string) string {

	//	bb := bytes.Buffer{}
	max := len(b)

	if len(a) > len(b) {
		max = len(a)
	}

	fmt.Println(max)

	arr := make([]int, max+1)

	for i := 1; i < len(arr); i++ {

		if len(a)-i >= 0 {
			arr[i] += int(a[len(a)-i]) - 48
		}
		if len(b)-i >= 0 {
			arr[i] += int(b[len(b)-i]) - 48
		}

	}
	fmt.Println(arr)

	for i := len(arr) - 1; i > 0; i-- {
		if arr[i] > 1 {
			arr[i] = 0
			arr[i-1]++
		}

	}

	//	fmt.Println(arr)

	return ""
}
