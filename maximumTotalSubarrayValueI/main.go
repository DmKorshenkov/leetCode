package main

import "math"

func main() {

}

func maxTotalValue(nums []int, k int) int64 {
	mimimum := math.MaxInt
	maximum := 0

	for _, n := range nums {
		maximum = max(n, maximum)
		minimum = min(n, minimum)
	}
	return int64(maximum-minimum) * int64(k)
}
