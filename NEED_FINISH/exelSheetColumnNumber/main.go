package main

import "fmt"

func main() {

	//	fmt.Println(titleToNumber("FXSHRXW"))
	fmt.Println(titleToNumber("AAA"))
}
func titleToNumber(columnTitle string) int {
	res := int(0)
	mapN := make([]int, len(columnTitle), len(columnTitle))
	for in := 0; in < len(columnTitle); in++ {
		mapN[in] = (int(columnTitle[in] - 64))
	}

	fmt.Println(mapN)
	return res
}
