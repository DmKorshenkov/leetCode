package main

import "fmt"

func main() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4}))
	fmt.Println(sortArrayByParity2([]int{3, 1, 2, 4}))
}

func sortArrayByParity(nums []int) []int {

	for in := 0; in < len(nums)-1; in++ {

		if nums[in]%2 != 0 {

			for next := in + 1; next < len(nums); next++ {
				if nums[next]%2 == 0 {
					nums[in], nums[next] = nums[next], nums[in]
					break
				}
			}
		}
	}
	return nums
}

func sortArrayByParity2(nums []int) []int {

	for i, j := 0, len(nums)-1; i < j; {

		if nums[i]%2 != 0 && nums[j]%2 == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		} else {
			if nums[i]%2 == 0 {
				i++
			}

			if nums[j]%2 != 0 {
				j--
			}
		}

	}

	return nums
}
