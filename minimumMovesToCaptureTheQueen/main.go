package main

func main() {

}

func minMovesToCaptureTheQueen(a int, b int, c int, d int, e int, f int) int {
	var s [9][9]int
	s[e][f] = 1 // queen
	s[a][b] = 2 // rook
	s[c][d] = 3 // bishop

	if rook(s, a, b) {
		return 1
	}
	if bishop(s, c, d) {
		return 1
	}
	return 2
}

func rook(table [9][9]int, x, y int) bool {
	// up && down
	for i, j := x, y; i <= 8; i++ {
		if table[i][j] == 1 {
			return true
		}
		if table[i][j] == 3 {
			break
		}
	}
	for i, j := x, y; i >= 1; i-- {
		if table[i][j] == 1 {
			return true
		}
		if table[i][j] == 3 {
			break
		}
	}
	// left && right
	for i, j := x, y; j >= 1; j-- {
		if table[i][j] == 1 {
			return true
		}
		if table[i][j] == 3 {
			break
		}
	}
	for i, j := x, y; j <= 8; j++ {
		if table[i][j] == 1 {
			return true
		}
		if table[i][j] == 3 {
			break
		}
	}
	return false
}
func bishop(table [9][9]int, x, y int) bool {
	// up+left && up+right
	for i, j := x, y; i >= 1 && j >= 1; {
		if table[i][j] == 1 {
			return true
		}
		if table[i][j] == 2 {
			break
		}
		i--
		j--
	}
	for i, j := x, y; i >= 1 && j <= 8; {
		if table[i][j] == 1 {
			return true
		}
		if table[i][j] == 2 {
			break
		}
		i--
		j++
	}
	// down+left && down+right
	for i, j := x, y; i <= 8 && j >= 1; {
		if table[i][j] == 1 {
			return true
		}
		if table[i][j] == 2 {
			break
		}
		i++
		j--
	}
	for i, j := x, y; i <= 8 && j <= 8; {
		if table[i][j] == 1 {
			return true
		}
		if table[i][j] == 2 {
			break
		}
		i++
		j++
	}
	return false
}
