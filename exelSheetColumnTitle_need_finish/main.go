package main

import "fmt"

func main() {
	fmt.Println(convertToTitle(26))
}

func convertToTitle(columnNumber int) string {
	str := string("")

	tmp := columnNumber / (26 + 1)

	fmt.Println(tmp)
	return str
}
