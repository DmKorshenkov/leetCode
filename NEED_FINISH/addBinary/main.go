package main

import (
	"fmt"
)

func main() {
	fmt.Println(addBinary("100", "0"))
	//	fmt.Println(int('1'), int('0'))
	//	fmt.Println(0 * degree(2, 1))
}

func addBinary(a string, b string) string {
	ln := len(a) - 1
	bin := int(0)
	for in := ln; in > -1; in-- {
		bin += int(a[in]-48) * degree(2, in-ln)
	}
	ln = len(b) - 1
	for in := ln; in > -1; in-- {
		bin += int(b[in]-48) * degree(2, in-ln)
	}
	return strBin(bin)
}

func strBin(num int) string {
	if num == 0 {
		return "0"
	}
	var slB = make([]string, 1, num)
	slB[0] = "0"
	fmt.Println(num)
	for n := 0; n < num; n++ {
		for in := 0; in < len(slB); in++ {
			if slB[in] == "0" {
				slB[in] = "1"
				break
			} else {
				slB[in] = "0"
				slB = append(slB, "1")
				break
			}
		}
	}

	///////
	///////
	//////
	var str string
	for in := len(slB) - 1; in > -1; in-- {
		str += slB[in]
	}
	return str
}

func degree(num int, degree int) int {
	if degree == 0 {
		return 1
	}
	for i := 1; i < degree; i++ {
		num *= 2
	}

	return num
}
