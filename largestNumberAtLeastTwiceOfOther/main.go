package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(dominantIndex([]int{3, 6, 1, 0}))
}

func dominantIndex(nums []int) int {
	ln := len(nums)
	duplicate := append(make([]int, 0, ln), nums...)
	sort.Ints(nums)
	if nums[ln-2]*2 <= nums[ln-1] {
		for in := range duplicate {
			if duplicate[in] == nums[ln-1] {
				return in
			}
		}
	}
	return -1
}
