package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(arrayPairSum([]int{1, 4, 3, 2}))
	fmt.Println(arrayPairSum([]int{1, 2}))
}

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	if len(nums) == 2 {
		return nums[0]
	}
	for i := 0; i < len(nums)-1; i += 2 {
		nums[0] += nums[i]
	}
	return nums[0]
}
