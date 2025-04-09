package main

import (
	"sort"
)

func main() {

}

func smallestRangeI(nums []int, k int) int {
	sort.Ints(nums)

	var res int = nums[len(nums)-1] - nums[0]
	var max int
	var min int
	for i := -k; i != k+1; i++ {
		min = nums[0] + i
		for j := -k; j != k+1; j++ {
			max = nums[len(nums)-1] + j
			if max-min < res {
				res = max - min
			}
		}
	}
	if res < 0 {
		return 0
	}
	return res
}
