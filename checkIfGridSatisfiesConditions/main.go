package main

import "fmt"

func main() {

	fmt.Println(satisfiesConditions([][]int{{2, 1, 3}, {2, 1, 3}, {2, 1, 0}, {2, 1, 3}, {2, 1, 3}}))
	fmt.Println(satisfiesConditions([][]int{{1, 1, 6, 1, 4, 6, 3, 1, 0, 7}}))
}

func satisfiesConditions(grid [][]int) bool {

	if len(grid) == 1 {
		for h := 0; h < len(grid[0])-1; h++ {
			if grid[0][h] == grid[0][h+1] {
				return false
			}
		}
	}

	for h := range grid[0] {

		for v := 0; v < len(grid)-1; v++ {

			if grid[v][h] != grid[v+1][h] {
				return false
			}
			if len(grid[v])-1 != h {
				if grid[v][h] == grid[v][h+1] {
					return false
				}
			}
		}

	}
	return true
}
