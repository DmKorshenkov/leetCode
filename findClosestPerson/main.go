func findClosest(x int, y int, z int) int {
    position1 := abc(x-z)
    position2 := abc(z-y)
    if position1 < position2 {
        return 1
    } 
    if position1 > position2 {
        return 2
    }
    return 0
}

func abc(n int) int {
    if n < 0 {
        return -n
    }
    return n
}
