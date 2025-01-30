package main

import "fmt"

func main() {
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{}({[]}){}"))
	// fmt.Println(int('('), int(')'))
	// fmt.Println(int('['), int(']'))
	// fmt.Println(int('{'), int('}'))
}

func isValid(s string) bool {
	return false
}

/*func isValid(s string) bool {
	var ok bool
	for now := 0; now < len(s); now++ {
		ok = false
		switch s[now] {
		case 41:
			ch = 40
		case 93:
			ch = 91
		case 125:
			ch = 123
		}
		for next := now; next > 0; next-- {
			if s[next] == ch {
				ok = true
				break
			}
		}
	}
	return ok
}*/
