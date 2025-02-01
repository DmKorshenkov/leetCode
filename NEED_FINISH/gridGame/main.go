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
	fmt.Println(grid[0])
	fmt.Println(grid[1])
	firstRobot(grid)
	fmt.Println(grid[0])
	fmt.Println(grid[1])
	secondRobot(grid)

	return secondRobot(grid)
}

func secondRobot(grid [][]int) int64 {
	var ln = len(grid[1]) - 1
	var sum int64
	for ln > -1 {
		sum += int64(grid[1][ln])
		if !(sumRight(grid[1][:ln]) > sumRight(grid[0][:ln])) {
			ln--
			break
		}
		ln--
	}

	for ln > -1 {
		sum += int64(grid[0][ln])
		ln--
	}
	return sum
}

func firstRobot(grid [][]int) {
	var in int
	for in < len(grid[0]) {
		if !(sumRight(grid[0][in:]) > sumRight(grid[1][in:])) {

			grid[0][in] = 0
			break
		}
		grid[0][in] = 0
		in++
	}

	for in < len(grid[1]) {
		grid[1][in] = 0
		in++
	}
	grid[1][len(grid[1])-1] = 0
}

func sumRight(arr []int) int64 {
	var sum int64
	for i := 0; i < len(arr); i++ {
		sum += int64(arr[i])
	}

	return sum
}
