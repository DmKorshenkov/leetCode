package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(containsDuplicate([]int{124, 231, 562, 1, 28, 29, 1, 829}))
}

func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for in := range nums {
		if in+1 == len(nums) {
			break
		}
		if nums[in] == nums[in+1] {
			return true
		}
	}
	return false
}
