package main

import "fmt"

func main() {

	fmt.Println(removeDuplicates([]int{1}))

}

func removeDuplicates(nums []int) int {
	var k = int(1)
	for now := 0; now < len(nums); {

		for next := now; next < len(nums); {
			if nums[now] < nums[next] {
				k++
				nums[now+1] = nums[next]
				now++
				continue
			}
			next++
		}
		return k
	}
	return k
}
