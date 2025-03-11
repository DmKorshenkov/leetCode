package main

import (
	"sort"
	"strconv"
)

func main() {

}

func findRelativeRanks(score []int) []string {
	dupl := make([]int, len(score))
	copy(dupl, score)
	step := 4
	sort.Ints(score)

	answer := make([]string, len(score))

	for i := len(score) - 1; i > -1; i-- {

		for j := range dupl {
			if dupl[j] == score[i] {
				if i == len(score)-1 {
					answer[j] = "Gold Medal"
					break
				}
				if i == len(score)-2 {
					answer[j] = "Silver Medal"
					break
				}
				if i == len(score)-3 {
					answer[j] = "Bronze Medal"
					break
				}
				answer[j] = strconv.Itoa(step)
				step++
				break
			}
		}
	}
	return answer
}
