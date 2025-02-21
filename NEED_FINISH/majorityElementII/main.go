package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(majorityElement([]int{1, 1, 1}))
}

func majorityElement(nums []int) []int {
	sort.Ints(nums)
	major_elems := make([]int, 0, len(nums)/3)
	for in := range nums {
		count := 1
		for next := in + 1; next < len(nums) && nums[in] == nums[next]; next++ {
			count++
		}
		if count > len(nums)/3 {
			major_elems = append(major_elems, nums[in])
		}
	}
	return major_elems
}
