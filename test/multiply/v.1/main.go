package main

import (
	"bytes"
	"fmt"
)

func main() {

	fmt.Println(multiply("123", "456"))

	//	fmt.Println(0 - 15)
	//
	// fmt.Println(999999999999 + 999)
}

func multiply(num1 string, num2 string) string {
	var res int
	//	var n1 int
	//	var n2 int
	var x2 = 1
	arr := make([]int, 0, 400)
	for i := len(num2) - 1; i > -1; i-- {
		x := 1
		//		n2 = (int(num2[i]) - 48) * x2
		for j := len(num1) - 1; j > -1; j-- {
			//			n1 = (int(num1[j]) - 48) * x
			fmt.Println((int(num1[j]) - 48) * x * (int(num2[i]) - 48) * x2)
			res = (int(num1[j]) - 48) * x * (int(num2[i]) - 48) * x2
			arr = append(arr, res)
			x *= 10
		}
		x2 *= 10

	}

	res = 0
	for i := range arr {
		res += arr[i]
	}
	fmt.Println(res)
	return ""
}

func multiply2(num1 string, num2 string) string {
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
