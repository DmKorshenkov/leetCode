package main

func main() {
	twoSum([]int{2, 3, 11, 15}, 9)
}

func twoSum(arr []int, target int) []int {
	for in := range arr {
		for i := in + 1; i < len(arr); i++ {
			if arr[in]+arr[i] == target {
				return []int{i, in}
			}
		}
	}
	return nil
}
