package main

func main() {

}

func maximumWealth(accounts [][]int) int {
	var wealth int
	var m, n = len(accounts), len(accounts[0])

	for i := 0; i < m; i++ {
		tmpWealth := 0
		for j := 0; j < n; j++ {
			tmpWealth += accounts[i][j]
		}
		wealth = max(wealth, tmpWealth)
	}
	return wealth
}
