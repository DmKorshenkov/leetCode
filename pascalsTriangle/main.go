package main

import "fmt"

func main() {
	fmt.Println(generate(5))
}

func generate(numRows int) [][]int {
	p := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		p[i] = make([]int, i+1)

		for j := 0; j < len(p[i]); j++ {

			if j == 0 || j == len(p[i])-1 {
				p[i][j] = 1
			} else {
				p[i][j] = p[i-1][j-1] + p[i-1][j]
			}
		}
	}
	return p
}
