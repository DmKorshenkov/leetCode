package main

import "fmt"

func main() {
	leftRightDifference([]int{10, 4, 8, 3})
	leftRightDifference([]int{1})

}

func leftRightDifference(nums []int) []int {

	arr := make([]int, len(nums))
	var sumL int
	var sumR int
	//	var diff int

	for i := 0; i < len(nums); i++ {
		sumR = 0
		for j := len(nums) - 1; j > i; j-- {
			sumR += nums[j]
		}

		arr[i] = sumL - sumR
		if sumL < sumR {
			arr[i] = sumR - sumL
		}
		//		fmt.Println(sumL, sumR, diff)
		sumL += nums[i]
	}
	fmt.Println(arr)
	//15 1 11 22
	return arr
}
