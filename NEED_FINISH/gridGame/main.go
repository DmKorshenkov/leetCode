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

	for in := range grid {
		fmt.Println(grid[in])
	}

	return 0
}

func sumHelp(num []int) (sum int64) {
	for i := range num {
		sum += int64(num[i])
	}
	return sum
}
