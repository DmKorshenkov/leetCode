package main

func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})
}

func moveZeroes(nums []int) {
	for in := range nums {
		for i := in + 1; i < len(nums); {
			if nums[in] != 0 {
				break
			} else {
				nums[i], nums[in] = nums[in], nums[i]
			}
			i++
		}
	}
}
