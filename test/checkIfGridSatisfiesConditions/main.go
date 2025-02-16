package main

import "fmt"

func main() {

	//	fmt.Println(satisfiesConditions([][]int{{1, 0, 2}, {1, 0, 2}}))
	//	fmt.Println(satisfiesConditions([][]int{{1, 3, 2}, {1, 3, 2}}))
	fmt.Println(satisfiesConditions([][]int{{1, 2}, {1, 2}, {1, 2}, {1, 3}}))
}

func satisfiesConditions(grid [][]int) bool {
	//	fmt.Println(grid)
	/*	for in := range grid {
		fmt.Println(grid[in])
	}*/
	h := 0
	ok := true
	for v := 0; v < len(grid); {

		if grid[v][h] != grid[v+1][h] && !(grid[v][h] == grid[v][h+1]) {
			ok = false
			break
		}

		v++
	}
	return ok
}
