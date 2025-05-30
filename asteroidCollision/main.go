func asteroidCollision(asteroids []int) []int {
	stack := []int{}

	for _, a := range asteroids {
		if len(stack) == 0 {
			stack = append(stack, a)
			continue
		}

		for len(stack) != 0 {
			last := stack[len(stack)-1]
			if a < 0 && last > 0 {
				if -a > last {
					stack = stack[:len(stack)-1]
					if len(stack) == 0 {
						stack = append(stack, a)
						break
					}

				} else if -a == last {
					stack = stack[:len(stack)-1]
					break
				} else {
					break
				}
			} else {
				stack = append(stack, a)
				break
			}

		}
	}
	return stack
}




func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0, len(asteroids))

	stack = append(stack, asteroids[0])

	for i := 1; i < len(asteroids); i++ {

		if asteroids[i] > 0 {
			stack = append(stack, asteroids[i])
			continue
		}
		fmt.Println(stack, asteroids[i])
		stack = collide(stack, Abs(asteroids[i]))
	}

	fmt.Println(stack)
	return stack
}

func collide(stack []int, a int) []int {
    if len(stack) == 0 {
        stack = append(stack, -a)
        return stack
    }
	if stack[len(stack)-1] < 0 {
		stack = append(stack, -a)
		return stack
	}

	if stack[len(stack)-1]-a == 0 {
		//		fmt.Println("!2")
		stack = stack[:len(stack)-1]

		return stack
	}

	if stack[len(stack)-1] > a {
		return stack
	}
	//	fmt.Println("!")
	stack = stack[:len(stack)-1]
	//	fmt.Println(stack)
	return collide(stack, a)
}

func Abs(n int) int {
    if n < 0 {
        return -n
    }

    return n
}
