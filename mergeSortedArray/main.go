package main

func main() {
	merge([]int{4, 0, 0, 0, 0, 0}, 1, []int{1, 2, 3, 5, 6}, 5)

}
func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	} else if m == 0 {
		for in := 0; in < n; in++ {
			nums1[in] = nums2[in]
		}
		return
	}
	for in, nn := len(nums1)-n, n-1; in < len(nums1); in++ {
		nums1[in] = nums2[nn]
		nn--
	}

	for now := 0; now < m+n; now++ {
		for next := 0; next < m+n; next++ {
			if nums1[next] > nums1[now] {
				nums1[now], nums1[next] = nums1[next], nums1[now]
			}
		}
	}

}
