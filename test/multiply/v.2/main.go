package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(multiply("498828660196", "840477629533"))
	//	fmt.Println(102 * 9)
	//
	// fmt.Println(byte(9))
}

func multiply(num1 string, num2 string) string {
	//	bb := bytes.Buffer{}
	mp := make(map[int]int)
	var n int
	var numbers int

	for i := 0; i < len(num1); i++ {
		n *= 10
		n += int(num1[i]) - 48
	}
	fmt.Println("num1 = ", n)
	for i := 0; i < 10; i++ {
		mp[i] = n * i
	}
	fmt.Println("\nmap : ")
	for key, val := range mp {
		fmt.Println("\nkey - ", key, "\nval - ", val)
	}
	fmt.Println()
	n = 1
	//	digits := []byte{}
	slice := make([]string, len(num2))
	for i := len(num2) - 1; i > -1; i-- {
		numbers = mp[int(num2[i])-48] * n
		digits := []byte{}
		for numbers > 0 {
			digit := numbers % 10
			digits = append([]byte{byte('0' + digit)}, digits...)
			numbers /= 10
		}
		slice[i] = string(digits)
		n *= 10
	}
	fmt.Println("\n slice :")

	for i := range slice {
		if slice[i] != string("") {
			fmt.Println("\nindex - ", i, "\nval - ", slice[i])
		}
	}
	var str string
	for i := range slice {
		if slice[i] != string("") {
			fmt.Println("\nstring -", str, "\n     slice -  ", slice[i])
			str = addString(str, slice[i])
		}
		//	str = addString(str, slice[i])

		//		fmt.Println("\nindex - ", i, "\nval -", slice[i])
	}
	//	fmt.Println("\nnumbers = ", numbers)

	return str
	//419254329864656431168468
	//8319439965581059540
	//4513610241269891072
}

func addString(num1 string, num2 string) string {
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

/*
	for i := 1; i < 10; i++ {
		numbers := n * i
		digits := []byte{}
		fmt.Println(len(digits), cap(digits))
		for numbers > 0 {
			digit := numbers % 10
			digits = append([]byte{byte('0' + digit)}, digits...)
			numbers /= 10

		}
		fmt.Println(len(digits), cap(digits))
		mp[byte(i)] = string(digits)
		digits = nil
		//		mp[byte(i)] =
	}
*/
