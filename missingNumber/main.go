package main

func main() {

}

func missingNumber(nums []int) int {
	sum, sumIn := int(0), len(nums)
	for in := range nums {
		sum = sum + nums[in]
		sumIn = sumIn + in
	}

	return sumIn - sum
}
