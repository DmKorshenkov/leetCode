package main

import "fmt"

func main() {
	fmt.Println(pivotIndex([]int{2, 1, -1}))
	fmt.Println(0 % 2)
}

func pivotIndex(nums []int) int {
	var sumR int

	for i := 0; i < len(nums); i++ {

		sumR += nums[i]
	}

	return -1
}
