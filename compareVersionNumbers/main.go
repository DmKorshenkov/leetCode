func compareVersion(version1 string, version2 string) int {

	nums1 := conversion(version1)
	nums2 := conversion(version2)
	if len(nums1) < len(nums2) {
		for i := len(nums1) - 1; i < len(nums2); i++ {
			nums1 = append(nums1, 0)
		}
	}
	if len(nums1) > len(nums2) {
		for i := len(nums2) - 1; i < len(nums1); i++ {
			nums2 = append(nums2, 0)
		}
	}
	for i := range nums1 {
		if nums1[i] < nums2[i] {
			return -1
		}
		if nums1[i] > nums2[i] {
			return 1
		}
	}
	return 0
}

func conversion(str string) []int {
	var nums []int
	left, right := 0, 0
	for right < len(str) {
		if str[right] == '.' {
			num, _ := strconv.Atoi(str[left:right])
			nums = append(nums, num)
			left = right + 1

		}
		right++
	}
	num, _ := strconv.Atoi(str[left:right])
	nums = append(nums, num)
	return nums
}
