package main

import "fmt"

func main() {
	//	fmt.Println(test([][]int{{2,5,4},{1,5,1}}))
	//	test([][]int{{3, 3, 1}, {8, 5, 2}})

	fmt.Println(gridGame([][]int{
		{1, 3, 7, 2, 2},
		{2, 1, 7, 2, 2}}))

}

func gridGame(grid [][]int) int64 {
	//fmt.Println("input: \n", grid[0], "\n", grid[1])
	firstRobot(grid)
	secondRobot(grid)
	//fmt.Println("output: \n", grid[0], "\n", grid[1])
	return secondRobot(grid)
}

func secondRobot(grid [][]int) int64 {
	var in int
	var sum int64
	for in < len(grid[1]) {
		if !(sumRight(grid[0][in:]) > sumRight(grid[1][in:])) {
			sum += int64(grid[1][in])
			break
		}
		sum += int64(grid[1][in])
		in++
	}

	for in < len(grid[1]) {
		sum += int64(grid[1][in])
		in++
	}
	return sum
}

func firstRobot(grid [][]int) {
	var in int
	//  	fmt.Println("input: \n", grid[0], "\n", grid[1])
	for in < len(grid[0]) {
		if !(sumRight(grid[0][in:]) > sumRight(grid[1][in:])) {
			//fmt.Println("down to - ", in)

			grid[0][in] = 0
			break
		}
		grid[0][in] = 0
		in++
	}
	//	fmt.Println("first end: \n", grid[0], "\n", grid[1])

	for in < len(grid[1]) {
		grid[1][in] = 0
		in++
	}
	grid[1][len(grid[1])-1] = 0
	//
	// fmt.Println("second end: \n", grid[0], "\n", grid[1])
}

func sumRight(arr []int) int64 {
	var sum int64
	for i := 0; i < len(arr); i++ {
		sum += int64(arr[i])
	}
	return sum
}
