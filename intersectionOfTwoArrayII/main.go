package main

import "fmt"

func main() {
	fmt.Println(intersect([]int{3, 1, 2}, []int{1, 1}))
}

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) < len(nums2) {
		array := make([]int, 0, len(nums1))
		for i := range nums1 {
			for j := range nums2 {
				if nums1[i] == nums2[j] {
					array = append(array, nums1[i])
					nums2[j] = -1
					break
				}
			}
		}
		return array
	}
	array := make([]int, 0, len(nums2))
	for i := range nums2 {
		for j := range nums1 {
			if nums2[i] == nums1[j] {
				array = append(array, nums2[i])
				nums1[j] = -1
				break
			}
		}
	}
	return array
}
