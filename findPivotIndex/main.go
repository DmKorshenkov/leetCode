package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Println(pivotIndex([]int{0}))
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6})) //28
}

func pivotIndex(nums []int) int {
	var sum1 int

	for i := 0; i < len(nums); i++ {

		sum1 += nums[i]
	}
	var sum2 int
	for i := 0; i < len(nums); i++ {
		sum1 -= nums[i]
		if sum1 == sum2 {
			return i
		}
		sum2 += nums[i]
	}

	return -1
}
