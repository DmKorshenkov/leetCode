package main

import "fmt"

func main() {
	fmt.Println(string(findTheDifference("mjcftlm", "mjcftlmy")))
}

func findTheDifference(s string, t string) byte {
	r1 := int(0)
	for in := range s {
		r1 += int(t[in])
		r1 -= int(s[in])
	}

	return byte(uint8(r1 + int(t[len(t)-1])))
}
