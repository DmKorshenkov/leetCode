package main

func main() {

}

func queensAttacktheKing(queens [][]int, king []int) [][]int {
	var res [][]int
	x, y := king[0], king[1]

	var S [8][8]int

	S[x][y] = 1
	//create table
	for i := range queens {
		S[queens[i][0]][queens[i][1]] = 2
	}
	//
	// up+left && up && up+right
	//
	for i, j := x, y; i >= 0 && j >= 0; {
		if S[i][j] == 2 {
			res = append(res, []int{i, j})
			break
		}
		i--
		j--
	}
	for i := x; i >= 0; {
		if S[i][y] == 2 {
			res = append(res, []int{i, y})
			break
		}
		i--
	}
	for i, j := x, y; i < 8 && j < 8; {
		if S[i][j] == 2 {
			res = append(res, []int{i, j})
			break
		}
		i++
		j++
	}
	//
	// mid+left && mid+right
	//
	for j := y; j >= 0; {
		if S[x][j] == 2 {
			res = append(res, []int{x, j})
			break
		}
		j--
	}
	for j := y; j < 8; {

		if S[x][j] == 2 {
			fmt.Println("?")
			res = append(res, []int{x, j})
			break
		}
		j++
	}
	//
	// down+left && down+right && down
	//
	for i, j := x, y; i < 8 && j >= 0; {
		if S[i][j] == 2 {
			res = append(res, []int{i, j})
			break
		}
		i++
		j--
	}
	for i, j := x, y; i >= 0 && j < 8; {
		if S[i][j] == 2 {
			res = append(res, []int{i, j})
			break
		}
		i--
		j++
	}
	for i := x; i < 8; {
		if S[i][y] == 2 {
			res = append(res, []int{i, y})
			break
		}
		i++
	}
	return res
}
