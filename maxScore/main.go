package main

import "fmt"

func main() {

	fmt.Println(maxScore("011101"))

}

func maxScore(s string) int {

	var scoreL int
	var scoreR int
	var res int

	for now := 0; now < len(s)-1; now++ {
		for next := now + 1; next < len(s); next++ {
			if s[next] == '1' {
				scoreR++
			}
		}
		if s[now] == '0' {
			scoreL++
		}
		if scoreL+scoreR > res {
			res = scoreL + scoreR
		}
		scoreR = 0
	}
	return res
}
