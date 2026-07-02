func evalRPN(tokens []string) int {
	ln := 0
	stack := []int{}
	for i := 0; i < len(tokens); i++ {
		ln = len(stack)
	//	fmt.Println(stack)
		switch tokens[i] {
		case "*":
			o1 := stack[ln-1]
			o2 := stack[ln-2]
			stack = stack[:ln-2]
			stack = append(stack, o1*o2)
		case "/":
			o1 := stack[ln-1]
			o2 := stack[ln-2]
			stack = stack[:ln-2]
			stack = append(stack, o2/o1)
		case "+":
			o1 := stack[ln-1]
			o2 := stack[ln-2]
			stack = stack[:ln-2]
			stack = append(stack, o1+o2)
		case "-":
			o1 := stack[ln-1]
			o2 := stack[ln-2]
			stack = stack[:ln-2]
			stack = append(stack, o2-o1)
		default:
			stack = append(stack, atoi(tokens[i]))
		}
	}
	return stack[len(stack)-1]
}

func atoi(str string) int {
	n, _ := strconv.Atoi(str)
	return n
}
