package main

import (
	"fmt"
)

func main() {
	i := gridGame([][]int{
		{2, 5, 4},
		{1, 5, 1}})
	fmt.Println(i)
}

func gridGame(grid [][]int) int64 {
	//	v := 0
	sum := int64(0)
	last := grid[1][len(grid[1])-1]
r1:
	for in := range grid[0] {
		if sumHelp(grid[0][in:])+int64(last) < sumHelp(grid[1]) {
			grid[0][in] = 0
			for in < len(grid[1]) {
				grid[1][in] = 0
				if in+1 == len(grid[1]) {
					break r1
				}
				in++
			}
		}
		grid[0][in] = 0
	}

	for in := range grid {
		fmt.Println(grid[in])
	}

	return sum
}

func sumHelp(num []int) (sum int64) {
	for i := range num {
		sum += int64(num[i])
	}
	return sum
}
