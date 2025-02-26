package main

import "fmt"

func main() {
	fmt.Println(isGood([]int{1, 2}))
}

func isGood(nums []int) bool {
	ln := len(nums)
	for in := 0; in < len(nums)-1; in++ {
		for next := in + 1; next < len(nums); next++ {
			if nums[in] > nums[next] {
				nums[in], nums[next] = nums[next], nums[in]
			}
		}
		if nums[in] != in+1 {
			return false
		}
	}
	if len(nums) >= 2 && nums[ln-1]+1 == len(nums) && nums[ln-1] == nums[ln-2] {
		return true
	}
	return false
}
