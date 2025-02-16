package main

import "fmt"

func main() {
	fmt.Println(byte('['))
	fmt.Println(isValid("[]"))
	// fmt.Println(isValid("((()))"))
}

func isValid(s string) bool {
	var st stack
	for in := range s {
		if s[in] == 40 || s[in] == 123 || s[in] == 91 {
			st.push(s[in])
		} else {
			if !st.peek(s[in]) {
				return false
			}
		}
	}
	return len(st.items) == 0
}

type stack struct {
	items []byte
}

func (s *stack) push(b byte) {
	s.items = append(s.items, b)
}

func (s *stack) peek(b byte) (ok bool) {

	if len(s.items) == 0 {
		return false
	}
	if (s.items[len(s.items)-1]+1 == b) ||
		(s.items[len(s.items)-1]+2 == b) {
		s.items = s.items[:len(s.items)-1]
		ok = true
	}
	return ok
}
