package main

import "fmt"

func main() {
	fmt.Println(sortArrayByParityII([]int{3, 2}))
}

func sortArrayByParityII(nums []int) []int {

	for in := 0; in < len(nums)-1; in++ {

		if in%2 == 0 && nums[in]%2 != 0 {
			for next := in + 1; next < len(nums); next++ {
				if nums[next]%2 == 0 {
					nums[in], nums[next] = nums[next], nums[in]
					break
				}
			}
		}
		if in%2 != 0 && nums[in]%2 == 0 {
			for next := in + 1; next < len(nums); next++ {
				if nums[next]%2 != 0 {
					nums[in], nums[next] = nums[next], nums[in]
					break
				}
			}
		}
	}
	return nums
}
