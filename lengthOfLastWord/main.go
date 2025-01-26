package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLastWord("hello"))
}

func lengthOfLastWord(s string) int {
	sl := strings.Split(s, " ")
	for in := len(sl) - 1; in > -1; in-- {
		//	fmt.Println(sl[in])
		if sl[in] != "" {
			return len(sl[in])
		}
	}
	return 0
}
