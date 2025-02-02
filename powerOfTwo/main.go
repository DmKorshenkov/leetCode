package main

import "fmt"

func main() {

	fmt.Println(isPowerOfTwo(1073741824))
	// fmt.Println(degree(-10))
}

func isPowerOfTwo(n int) bool {
	if n == 1 {
		return true
	}
	num := 2
	for x := 2; num <= n; x++ {
		if num == n {
			return true
		}
		num *= 2
	}

	return false
}
