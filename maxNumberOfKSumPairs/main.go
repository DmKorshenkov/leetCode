package main

func maxOperations(nums []int, k int) int {

	var res int

	res = 0

	for i := 0; i < len(nums); i++ {
		x := k - nums[i]
		for j := i + 1; x > 0 && j < len(nums); j++ {
			if nums[j] == x {
				nums = remove(nums, j)
				res++
				break
			}
		}
	}

	return res
}

func remove(nums []int, i int) []int {
	nums[i] = nums[len(nums)-1]
	return nums[:len(nums)-1]
}
