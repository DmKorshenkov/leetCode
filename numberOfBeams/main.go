package main

func main() {
	//	numberOfBeams([]string{"011001","000000","010100","001000"})
	numberOfBeams([]string{"011001", "000000", "010100"})

}

func numberOfBeams(bank []string) int {
	var a int
	var now int
	var next int

	for i := range bank {
		now = 0
		for j := range bank[i] {
			if bank[i][j] == '1' {
				now++
			}
		}
		if now != 0 && next != 0 {
			a += next * now
		}
		if now != 0 {
			next = now
		}
	}

	//	fmt.Println(a)
	return a
}
