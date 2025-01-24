package main

import "fmt"

func main() {
	fmt.Println(removeElement([]int{3, 2}, 3))
}
func removeElement(nums []int, val int) int {
	var k int
	for now := 0; now < len(nums); {
		if nums[now] == val {
			for last := len(nums) - 1; last > now; {
				if nums[last] != val {
					nums[now], nums[last] = nums[last], nums[now]
					break
				}
				last--
			}
		}
		now++
	}

	for ln := len(nums) - 1; ln > -1; ln-- {
		if nums[ln] != val {
			k++
		}
	}

	return k
}
