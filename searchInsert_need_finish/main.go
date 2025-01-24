package main

import "fmt"

func main() {

	fmt.Println(searchInsert([]int{1, 3, 5}, 2))
}

func searchInsert(arr []int, target int) int {

	low, hight := 0, len(arr)-1
	var mid int
	for low <= hight {

		mid = low + (hight-low)/2
		fmt.Printf("low: - %v	hight: - %v		mid:- %v\n", low, hight, mid)

		if arr[mid] == target {
			return mid
		}

		if target > arr[mid] {
			low = mid + 1
		}

		if target < arr[mid] {
			hight = mid - 1
		}

	}
	if arr[mid] < target && arr[mid+1] > target {
		return mid + 1
	}
	return mid
}
