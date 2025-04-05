package main

func main() {
	minOperations("001")
}

func minOperations(boxes string) []int {
	n := len(boxes)
	array := make([]int, len(boxes))
	mp := make(map[int]int)
	for ball := range boxes {
		if boxes[ball] == '1' {
			mp[ball] = 1
		} else {
			mp[ball] = 0
		}
	}

	for key := range mp {
		step := 0
		if mp[key] == 1 {
			for i := key + 1; i < n; i++ {
				step++
				array[i] += step
			}
		}
	}

	for key := n - 1; key > -1; key-- {
		step := 0
		if mp[key] == 1 {
			for i := key - 1; i > -1; i-- {
				step++
				array[i] += step
			}
		}
	}

	return array
}
