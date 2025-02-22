package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(3 % 3)
}

func fizzBuzz(n int) []string {
	answers := make([]string, n)
	for i := 1; i <= n; i++ {
		switch i % 15 {
		case 0:
			answers[i-1] = "FizzBuzz"
		case 3, 6, 9, 12:
			answers[i-1] = "Fizz"
		case 5, 10:
			answers[i-1] = "Buzz"
		default:
			answers[i-1] = strconv.Itoa(i)
		}
	}
	return answers
}
