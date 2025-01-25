package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 6))

}

func searchInsert(arr []int, target int) int {
	low, hight := 0, len(arr)-1
	if target > arr[hight] {
		return len(arr)
	} else if target < arr[low] {
		return 0
	} else if target == arr[hight] {
		return hight
	}

	if hight == 1 {
		if target == arr[0] {
			return 0
		}
		return 1

	}
	for low <= hight {
		mid := low + (hight-low)/2
		//	fmt.Printf("arr[mid] = %v\nlow: %v mid: %v hight: %v\n\n", arr[mid], low, mid, hight)
		if arr[mid] == target {
			//		fmt.Printf("target = %v\nmid: %v\n", target, mid)
			return mid
		}

		if arr[mid] > target && arr[mid-1] < target {
			//		fmt.Println("!!if - ", mid)
			return mid
		}

		if arr[mid] < target && arr[mid+1] > target {
			return mid + 1
		}

		if arr[mid] > target {
			hight = mid
		}

		if arr[mid] < target {
			low = mid
		}

	}
	return 0
}
